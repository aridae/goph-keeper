package secretsapi

import (
	"context"
	"fmt"
	secretusecases "github.com/aridae/goph-keeper/internal/server/controllers/secret"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
)

func (i *Implementation) CreateSecret(ctx context.Context, req *desc.CreateSecretRequest) (*desc.CreateSecretResponse, error) {
	err := i.useCasesController.CreateSecret(ctx, secretusecases.CreateSecretRequest{
		Key:  req.GetKey(),
		Data: req.GetData(),
		Meta: req.GetMeta(),
	})
	if err != nil {
		return nil, fmt.Errorf("useCasesController.CreateSecret: %w", err)
	}

	return &desc.CreateSecretResponse{}, nil
}
