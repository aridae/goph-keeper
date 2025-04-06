package secret

import (
	"context"
	"errors"
	"fmt"
	"github.com/aridae/goph-keeper/internal/common/domain-errors"
	"github.com/aridae/goph-keeper/internal/server/auth/authctx"
	"github.com/aridae/goph-keeper/internal/server/models"
	secretrepo "github.com/aridae/goph-keeper/internal/server/repos/secret"
)

type CreateSecretRequest struct {
	Key  string
	Data []byte
	Meta map[string]string
}

func (c *Controller) CreateSecret(ctx context.Context, req CreateSecretRequest) error {
	user, isAuthorized := authctx.GetUserFromContext(ctx)
	if !isAuthorized {
		return domainerrors.ErrUnauthorized()
	}

	secret := models.Secret{
		Accessor: models.SecretAccessor{
			OwnerUsername: user.Username,
			Key:           req.Key,
		},
		Data: req.Data,
		Meta: req.Meta,
	}

	err := c.secretRepository.CreateSecret(ctx, secret, c.now())
	if err != nil {
		if errors.Is(err, secretrepo.ErrOwnerUsernameKeyCombinationConstraintViolated) {
			return domainerrors.ErrSecretAlreadyExists(req.Key)
		}

		return fmt.Errorf("secretRepository.CreateSecret: %w", err)
	}

	return nil
}
