package secretsapi

import (
	"context"
	secretusecases "github.com/aridae/goph-keeper/internal/server/controllers/secret"
	"github.com/aridae/goph-keeper/internal/server/models"
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
