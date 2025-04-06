package secretsservice

import (
	"context"
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
)

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
