package user

import (
	"context"
	"fmt"
	domainerrors "github.com/aridae/goph-keeper/internal/server/models/domain-errors"
	"github.com/aridae/goph-keeper/internal/server/pkg/jwt"
)

type LoginUserRequest struct {
	Username string
	Password string
}

type LoginUserResponse struct {
	JWT string
}

func (c *Controller) LoginUser(ctx context.Context, req LoginUserRequest) (LoginUserResponse, error) {
	user, err := c.userRepository.GetByUsername(ctx, req.Username)
	if err != nil {
		return LoginUserResponse{}, fmt.Errorf("error getting user credentials: %w", err)
	}

	if user == nil {
		return LoginUserResponse{}, domainerrors.ErrInvalidUserCredentials()
	}

	if !user.Equal(req.Username, req.Password) {
		return LoginUserResponse{}, domainerrors.ErrInvalidUserCredentials()
	}
	token, err := c.jwtService.GenerateToken(ctx, jwt.Claims{Subject: user.Username})
	if err != nil {
		return LoginUserResponse{}, fmt.Errorf("failed to create JWT token: %w", err)
	}

	return LoginUserResponse{JWT: token}, nil
}
