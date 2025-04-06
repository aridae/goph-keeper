package usersapi

import (
	"context"
	"github.com/aridae/goph-keeper/internal/server/usecases/user"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
)

type useCasesController interface {
	LoginUser(ctx context.Context, req user.LoginUserRequest) (user.LoginUserResponse, error)
	RegisterUser(ctx context.Context, req user.RegisterUserRequest) (user.RegisterUserResponse, error)
}

type Implementation struct {
	desc.UnimplementedUsersServiceServer

	useCasesController useCasesController
}

func New(useCasesController useCasesController) *Implementation {
	return &Implementation{useCasesController: useCasesController}
}
