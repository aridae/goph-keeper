package usersservice

import (
	"context"
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
)

// RegisterUser регистрирует нового пользователя в сервисе UsersService.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// creds UserCredentials - учетные данные пользователя для регистрации.
//
// Возвращаемые значения:
// string - токен аутентификации, если регистрация успешна.
// error - ошибка, если не удалось зарегистрировать пользователя.
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
