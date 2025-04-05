package main

import (
	"github.com/aridae/goph-keeper/internal/client/config"
	secretsservice "github.com/aridae/goph-keeper/internal/client/downstream/secrets-service"
	"github.com/aridae/goph-keeper/internal/client/prompt"
	createsecret "github.com/aridae/goph-keeper/internal/client/usecases/create-secret"
	"github.com/aridae/goph-keeper/internal/logger"
	"github.com/spf13/cobra"
)

func main() {
	cnf := config.MustGetOnce()

	//usersServiceClient, err := usersservice.NewClient(cnf.UsersServiceHost)
	//if err != nil {
	//	logger.Fatalf("failed to create users service client: %v", err)
	//}

	secretsServiceClient, err := secretsservice.NewClient(cnf.SecretsServiceHost)
	if err != nil {
		logger.Fatalf("failed to create secrets service client: %v", err)
	}

	createSecretUseCase := createsecret.NewHandler(secretsServiceClient)

	promptIO := prompt.NewService(
		createSecretUseCase,
	)

	root := mustRegisterCommands(promptIO)

	if err := root.Execute(); err != nil {
		logger.Fatalf("failed to execute root command: %v", err)
	}
}

func mustRegisterCommands(promptIO *prompt.Service) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "goph-keeper",
		Short: "Use goph-keeper to keep your secrets safe and sound",
		Long:  `Manage your secrets safely with the goph-keeper CLI app by your side`,
	}

	createSecretCommand := &cobra.Command{
		Use:   "new",
		Short: "Creates a new secret record",
		Long:  `Creates a new  secret record`,
		Run: func(cmd *cobra.Command, args []string) {
			promptIO.CreateSecret()
		},
	}
	rootCmd.AddCommand(createSecretCommand)

	return rootCmd
}
