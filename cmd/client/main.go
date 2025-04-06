package main

import (
	"context"
	"github.com/abiosoft/ishell/v2"
	"github.com/aridae/goph-keeper/internal/client/auth"
	"github.com/aridae/goph-keeper/internal/client/config"
	authmw "github.com/aridae/goph-keeper/internal/client/downstream/grpc-client-mw/auth-mw"
	errormappingmw "github.com/aridae/goph-keeper/internal/client/downstream/grpc-client-mw/error-mapping-mw"
	secretsservice "github.com/aridae/goph-keeper/internal/client/downstream/secrets-service"
	usersservice "github.com/aridae/goph-keeper/internal/client/downstream/users-service"
	"github.com/aridae/goph-keeper/internal/client/prompt"
	createsecret "github.com/aridae/goph-keeper/internal/client/usecases/create-secret"
	getsecret "github.com/aridae/goph-keeper/internal/client/usecases/get-secret"
	loginuser "github.com/aridae/goph-keeper/internal/client/usecases/login-user"
	registeruser "github.com/aridae/goph-keeper/internal/client/usecases/register-user"
	"github.com/aridae/goph-keeper/internal/common/logger"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	ctx := context.Background()
	logger.Infof("Starting goph-keeper client cli app with build flags:\nBuild version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)

	cnf := config.MustGetOnce()

	sessionStorage := auth.NewSession()

	usersServiceClient, err := usersservice.NewClient(
		cnf.UsersServiceHost,
		errormappingmw.ErrorMapperInterceptor(),
	)
	if err != nil {
		logger.Fatalf("failed to create users service client: %v", err)
	}

	secretsServiceClient, err := secretsservice.NewClient(
		cnf.SecretsServiceHost,
		errormappingmw.ErrorMapperInterceptor(),
		authmw.AuthInterceptor(sessionStorage),
	)
	if err != nil {
		logger.Fatalf("failed to create secrets service client: %v", err)
	}

	registerUserHandler := registeruser.NewHandler(usersServiceClient, sessionStorage)
	loginUserHandler := loginuser.NewHandler(usersServiceClient, sessionStorage)
	createSecretUseCase := createsecret.NewHandler(secretsServiceClient)
	getSecretUseCase := getsecret.NewHandler(secretsServiceClient)

	promptIO := prompt.NewService(
		registerUserHandler,
		loginUserHandler,
		createSecretUseCase,
		getSecretUseCase,
	)

	shell := mustRegisterShellCommands(ctx, promptIO)

	shell.Run()
}

func mustRegisterShellCommands(ctx context.Context, promptIO *prompt.Service) *ishell.Shell {
	shell := ishell.New()
	shell.AutoHelp(true)

	createSecretCommand := &ishell.Cmd{
		Name: "new-secret",
		Help: "Creates a new secret record",
		Func: func(c *ishell.Context) {
			promptIO.RunCreateSecretPrompt(ctx)
		},
	}
	shell.AddCmd(createSecretCommand)

	getSecretCommand := &ishell.Cmd{
		Name: "get-secret",
		Help: "Obtains a secret record",
		Func: func(c *ishell.Context) {
			promptIO.RunGetSecretPrompt(ctx)
		},
	}
	shell.AddCmd(getSecretCommand)

	loginUserCommand := &ishell.Cmd{
		Name: "login",
		Help: "Login in system",
		Func: func(c *ishell.Context) {
			promptIO.RunLoginUserPrompt(ctx)
		},
	}
	shell.AddCmd(loginUserCommand)

	registerUserCommand := &ishell.Cmd{
		Name: "register",
		Help: "Register in system",
		Func: func(c *ishell.Context) {
			promptIO.RunRegisterUserPrompt(ctx)
		},
	}
	shell.AddCmd(registerUserCommand)

	return shell
}
