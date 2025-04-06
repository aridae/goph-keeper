package secretsapi

import (
	"context"
	"fmt"
	"github.com/aridae/goph-keeper/internal/server/models"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) GetSecret(ctx context.Context, req *desc.GetSecretRequest) (*desc.GetSecretResponse, error) {
	secret, err := i.useCasesController.GetSecret(ctx, req.GetKey())
	if err != nil {
		return nil, fmt.Errorf("useCasesController.GetSecret: %w", err)
	}

	return &desc.GetSecretResponse{
		Secret: mapDomainToAPISecret(secret),
	}, nil
}

func mapDomainToAPISecret(secret models.Secret) *desc.Secret {
	return &desc.Secret{
		Key:       secret.Accessor.Key,
		Data:      secret.Data,
		Meta:      secret.Meta,
		CreatedAt: timestamppb.New(secret.CreatedAt),
		UpdatedAt: timestamppb.New(secret.UpdatedAt),
	}
}
