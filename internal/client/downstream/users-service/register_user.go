package usersservice

import (
	"context"
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
)

func (c *Client) RegisterUser(ctx context.Context, creds UserCredentials) (string, error) {
	resp, err := c.grpc.RegisterUser(ctx, &desc.RegisterUserRequest{
		Username: creds.Username,
		Password: creds.Password,
	})
	if err != nil {
		return "", fmt.Errorf("grpc.LoginUser: %w", err)
	}

	return resp.GetToken(), nil
}
