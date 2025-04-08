package secretsservice

import (
	"context"
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
)

// CreateSecret создаёт новый секрет в сервисе SecretsService.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// secret Secret - информация о секрете для создания.
//
// Возвращаемые значения:
// error - ошибка, если не удалось создать секрет.
func (c *Client) CreateSecret(ctx context.Context, secret Secret) error {
	_, err := c.grpc.CreateSecret(ctx, &desc.CreateSecretRequest{
		Key:  secret.Key,
		Data: secret.Data,
		Meta: secret.Meta,
	})
	if err != nil {
		return fmt.Errorf("grpc.CreateSecret: %w", err)
	}

	return nil
}
