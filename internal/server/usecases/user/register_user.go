package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/aridae/goph-keeper/internal/common/domain-errors"
	"github.com/aridae/goph-keeper/internal/server/models"
	"github.com/aridae/goph-keeper/internal/server/pkg/jwt"
	userrepo "github.com/aridae/goph-keeper/internal/server/repos/user"
)

type RegisterUserRequest struct {
	Username string
	Password string
}

type RegisterUserResponse struct {
	JWT string
}

func (c *Controller) RegisterUser(ctx context.Context, req RegisterUserRequest) (RegisterUserResponse, error) {
	now := c.now()

	creds, err := models.NewUserCredentials(req.Username, req.Password)
	if err != nil {
		return RegisterUserResponse{}, fmt.Errorf("failed to create user credentials: %w", err)
	}

	err = c.userRepository.CreateUser(ctx, creds, now)
	if err != nil {
		if errors.Is(err, userrepo.ErrUsernameUniqueConstraintViolated) {
			return RegisterUserResponse{}, domainerrors.ErrUsernameAlreadyTaken(req.Username)
		}

		return RegisterUserResponse{}, fmt.Errorf("error creating user: %w", err)
	}

	token, err := c.jwtService.GenerateToken(ctx, jwt.Claims{Subject: creds.Username})
	if err != nil {
		return RegisterUserResponse{}, fmt.Errorf("failed to create JWT token: %w", err)
	}

	return RegisterUserResponse{JWT: token}, nil
}
