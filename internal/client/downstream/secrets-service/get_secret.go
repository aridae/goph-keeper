package secretsservice

import (
	"context"
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
)

// GetSecret получает секрет из сервиса SecretsService по указанному ключу.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// key string - ключ секрета.
//
// Возвращаемые значения:
// Secret - информация о секрете.
// error - ошибка, если не удалось получить секрет.
func (c *Client) GetSecret(ctx context.Context, key string) (Secret, error) {
	resp, err := c.grpc.GetSecret(ctx, &desc.GetSecretRequest{Key: key})
	if err != nil {
		return Secret{}, fmt.Errorf("grpc.GetSecret: %w", err)
	}

	return Secret{
		Key:  resp.GetSecret().GetKey(),
		Data: resp.GetSecret().GetData(),
		Meta: resp.GetSecret().GetMeta(),
	}, nil
}
