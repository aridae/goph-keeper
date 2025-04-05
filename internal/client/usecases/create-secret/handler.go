package createsecret

import (
	"context"
	"fmt"
	secretsservice "github.com/aridae/goph-keeper/internal/client/downstream/secrets-service"
)

type secretsService interface {
	CreateSecret(ctx context.Context, secret secretsservice.Secret) error
}

type Handler struct {
	secretsService secretsService
}

func NewHandler(secretsService secretsService) *Handler {
	return &Handler{
		secretsService: secretsService,
	}
}

type Request struct {
	Key  string
	Data []byte
}

func (h *Handler) Handle(ctx context.Context, req Request) error {
	err := h.secretsService.CreateSecret(ctx, secretsservice.Secret{
		Key:  req.Key,
		Data: req.Data,
	})
	if err != nil {
		return fmt.Errorf("secretsService.CreateSecret: %w", err)
	}

	return nil
}
