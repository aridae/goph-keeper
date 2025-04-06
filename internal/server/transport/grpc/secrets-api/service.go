package secretsapi

import (
	"context"
	"github.com/aridae/goph-keeper/internal/server/models"
	secretusecases "github.com/aridae/goph-keeper/internal/server/usecases/secret"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
)

type useCasesController interface {
	CreateSecret(ctx context.Context, req secretusecases.CreateSecretRequest) error
	GetSecret(ctx context.Context, key string) (models.Secret, error)
}

type Implementation struct {
	desc.UnimplementedSecretsServiceServer

	useCasesController useCasesController
}

func New(useCasesController useCasesController) *Implementation {
	return &Implementation{useCasesController: useCasesController}
}
