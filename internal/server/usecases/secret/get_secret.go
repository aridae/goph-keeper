package secret

import (
	"context"
	"fmt"
	"github.com/aridae/goph-keeper/internal/common/domain-errors"
	"github.com/aridae/goph-keeper/internal/server/auth/authctx"
	"github.com/aridae/goph-keeper/internal/server/models"
)

func (c *Controller) GetSecret(ctx context.Context, key string) (models.Secret, error) {
	user, isAuthorized := authctx.GetUserFromContext(ctx)
	if !isAuthorized {
		return models.Secret{}, domainerrors.ErrUnauthorized()
	}

	secretAccessor := models.SecretAccessor{
		OwnerUsername: user.Username,
		Key:           key,
	}

	secret, err := c.secretRepository.GetByAccessor(ctx, secretAccessor)
	if err != nil {
		return models.Secret{}, fmt.Errorf("secretRepository.CreateSecret: %w", err)
	}

	if secret == nil {
		return models.Secret{}, domainerrors.ErrSecretNotFound(key)
	}

	return *secret, nil
}
