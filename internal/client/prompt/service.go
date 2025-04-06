package prompt

import (
	"context"
	"errors"
	secretsservice "github.com/aridae/goph-keeper/internal/client/downstream/secrets-service"
	createsecret "github.com/aridae/goph-keeper/internal/client/usecases/create-secret"
	loginuser "github.com/aridae/goph-keeper/internal/client/usecases/login-user"
	registeruser "github.com/aridae/goph-keeper/internal/client/usecases/register-user"
	domainerrors "github.com/aridae/goph-keeper/internal/common/domain-errors"
)

type registerUserHandler interface {
	Handle(ctx context.Context, req registeruser.Request) error
}

type loginUserHandler interface {
	Handle(ctx context.Context, req loginuser.Request) error
}

type createSecretHandler interface {
	Handle(ctx context.Context, req createsecret.Request) error
}

type getSecretHandler interface {
	Handle(ctx context.Context, key string) (secretsservice.Secret, error)
}

type Service struct {
	registerUserHandler registerUserHandler
	loginUserHandler    loginUserHandler
	createSecretHandler createSecretHandler
	getSecretHandler    getSecretHandler
}

func NewService(
	registerUserHandler registerUserHandler,
	loginUserHandler loginUserHandler,
	createSecretHandler createSecretHandler,
	getSecretHandler getSecretHandler,
) *Service {
	return &Service{
		registerUserHandler: registerUserHandler,
		loginUserHandler:    loginUserHandler,
		createSecretHandler: createSecretHandler,
		getSecretHandler:    getSecretHandler,
	}
}

func printableErrorMessage(err error) string {
	if domainErr := new(domainerrors.DomainError); errors.As(err, domainErr) {
		return domainErr.Msg
	}

	return err.Error()
}
