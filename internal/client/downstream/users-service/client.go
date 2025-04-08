package usersservice

import (
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client предоставляет доступ к сервису UsersService через gRPC.
type Client struct {
	conn *grpc.ClientConn
	grpc desc.UsersServiceClient
}

// NewClient создаёт новый клиент для взаимодействия с сервисом UsersService.
//
// Параметры:
// target string - адрес сервера gRPC.
// interceptors ...grpc.UnaryClientInterceptor - цепочка интерсепторов для обработки запросов.
//
// Возвращаемые значения:
// *Client - новый клиент UsersService.
// error - ошибка, если не удалось подключиться к серверу.
func NewClient(target string, interceptors ...grpc.UnaryClientInterceptor) (*Client, error) {
	cc, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptors...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial <target:%s>: %w", target, err)
	}

	return &Client{
		conn: cc,
		grpc: desc.NewUsersServiceClient(cc),
	}, nil
}

// Close закрывает соединение с сервисом UsersService.
//
// Возвращаемые значения:
// error - ошибка, если не удалось закрыть соединение.
func (c *Client) Close() error {
	return c.conn.Close()
}
