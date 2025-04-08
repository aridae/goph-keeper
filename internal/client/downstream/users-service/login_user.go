package usersservice

import (
	"context"
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
)

// LoginUser выполняет вход пользователя в сервис UsersService.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// creds UserCredentials - учетные данные пользователя для входа.
//
// Возвращаемые значения:
// string - токен аутентификации, если вход успешен.
// error - ошибка, если не удалось выполнить вход.
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
