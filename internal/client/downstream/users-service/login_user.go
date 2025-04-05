package usersservice

import (
	"context"
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
)

func (c *Client) LoginUser(ctx context.Context, creds UserCredentials) (string, error) {
	resp, err := c.grpc.LoginUser(ctx, &desc.LoginUserRequest{
		Username: creds.Username,
		Password: creds.Password,
	})
	if err != nil {
		return "", fmt.Errorf("grpc.LoginUser: %w", err)
	}

	return resp.GetToken(), nil
}
