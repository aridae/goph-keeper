package registeruser

import (
	"context"
	"fmt"
	usersservice "github.com/aridae/goph-keeper/internal/client/downstream/users-service"
)

type usersService interface {
	RegisterUser(ctx context.Context, creds usersservice.UserCredentials) (string, error)
}

type sessionStorage interface {
	StoreToken(ctx context.Context, token string) error
}

type Handler struct {
	usersService   usersService
	sessionStorage sessionStorage
}

func NewHandler(
	usersService usersService,
	sessionStorage sessionStorage,
) *Handler {
	return &Handler{
		usersService:   usersService,
		sessionStorage: sessionStorage,
	}
}

type Request struct {
	Login    string
	Password string
}

func (h *Handler) Handle(ctx context.Context, req Request) error {
	token, err := h.usersService.RegisterUser(ctx, usersservice.UserCredentials{
		Username: req.Login,
		Password: req.Password,
	})
	if err != nil {
		return fmt.Errorf("usersService.RegisterUser: %w", err)
	}

	err = h.sessionStorage.StoreToken(ctx, token)
	if err != nil {
		return fmt.Errorf("sessionStorage.StoreToken: %w", err)
	}

	return nil
}
