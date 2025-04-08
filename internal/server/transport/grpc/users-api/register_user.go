package usersapi

import (
	"context"
	"fmt"
	"github.com/aridae/goph-keeper/internal/server/usecases/user"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
)

func (i *Implementation) RegisterUser(ctx context.Context, req *desc.RegisterUserRequest) (*desc.RegisterUserResponse, error) {
	resp, err := i.useCasesController.RegisterUser(ctx, user.RegisterUserRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, fmt.Errorf("useCasesController.RegisterUser: %w", err)
	}

	return &desc.RegisterUserResponse{Token: resp.JWT}, nil
}
