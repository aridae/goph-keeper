package secret

import (
	"context"
	"github.com/aridae/goph-keeper/internal/server/models"
	"time"
)

type secretRepository interface {
	CreateSecret(ctx context.Context, secret models.Secret, now time.Time) error
	GetByAccessor(ctx context.Context, accessor models.SecretAccessor) (*models.Secret, error)
}

type Controller struct {
	secretRepository secretRepository
	now              func() time.Time
}

func NewController(
	secretRepository secretRepository,
) *Controller {
	return &Controller{
		secretRepository: secretRepository,
		now:              time.Now().UTC,
	}
}
