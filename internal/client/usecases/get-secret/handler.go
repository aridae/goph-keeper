package getsecret

import (
	"context"
	"fmt"
	secretsservice "github.com/aridae/goph-keeper/internal/client/downstream/secrets-service"
)

type secretsService interface {
	GetSecret(ctx context.Context, key string) (secretsservice.Secret, error)
}

type Handler struct {
	secretsService secretsService
}

func NewHandler(secretsService secretsService) *Handler {
	return &Handler{
		secretsService: secretsService,
	}
}

func (h *Handler) Handle(ctx context.Context, key string) (secretsservice.Secret, error) {
	secret, err := h.secretsService.GetSecret(ctx, key)
	if err != nil {
		return secretsservice.Secret{}, fmt.Errorf("secretsService.GetSecret: %w", err)
	}

	return secret, nil
}
