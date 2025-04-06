package user

import (
	"context"
	"github.com/aridae/goph-keeper/internal/server/models"
	"github.com/aridae/goph-keeper/internal/server/pkg/jwt"
	"time"
)

type userRepository interface {
	CreateUser(ctx context.Context, user models.UserCredentials, now time.Time) error
	GetByUsername(ctx context.Context, username string) (*models.UserCredentials, error)
}

type jwtService interface {
	GenerateToken(ctx context.Context, claims jwt.Claims) (string, error)
}

type Controller struct {
	userRepository userRepository
	jwtService     jwtService
	now            func() time.Time
}

func NewController(
	userRepository userRepository,
	jwtService jwtService,
) *Controller {
	return &Controller{
		userRepository: userRepository,
		jwtService:     jwtService,
		now:            time.Now().UTC,
	}
}
