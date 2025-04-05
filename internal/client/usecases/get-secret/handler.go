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

type Command struct {
	Key string
}

func (h *Handler) Handle(ctx context.Context, command *Command) (secretsservice.Secret, error) {
	secret, err := h.secretsService.GetSecret(ctx, command.Key)
	if err != nil {
		return secretsservice.Secret{}, fmt.Errorf("secretsService.CreateSecret: %w", err)
	}

	return secret, nil
}
