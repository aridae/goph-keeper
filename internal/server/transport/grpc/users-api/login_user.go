package usersapi

import (
	"context"
	"fmt"
	"github.com/aridae/goph-keeper/internal/server/usecases/user"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
)

func (i *Implementation) LoginUser(ctx context.Context, req *desc.LoginUserRequest) (*desc.LoginUserResponse, error) {
	resp, err := i.useCasesController.LoginUser(ctx, user.LoginUserRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, fmt.Errorf("useCasesController.LoginUser: %w", err)
	}

	return &desc.LoginUserResponse{Token: resp.JWT}, nil
}
