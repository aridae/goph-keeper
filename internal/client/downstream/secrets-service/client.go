package secretsservice

import (
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client предоставляет доступ к сервису SecretsService через gRPC.
type Client struct {
	conn *grpc.ClientConn
	grpc desc.SecretsServiceClient
}

// NewClient создаёт новый клиент для взаимодействия с сервисом SecretsService.
//
// Параметры:
// target string - адрес сервера gRPC.
// interceptors ...grpc.UnaryClientInterceptor - цепочка интерсепторов для обработки запросов.
//
// Возвращаемые значения:
// *Client - новый клиент SecretsService.
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
		grpc: desc.NewSecretsServiceClient(cc),
	}, nil
}

// Close закрывает соединение с сервисом SecretsService.
//
// Возвращаемые значения:
// error - ошибка, если не удалось закрыть соединение.
func (c *Client) Close() error {
	return c.conn.Close()
}
