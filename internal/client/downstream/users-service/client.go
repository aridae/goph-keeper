package usersservice

import (
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	grpc desc.UsersServiceClient
}

// NewClient создает клиент.
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

// Close закрывает соединение с сервисом.
func (c *Client) Close() error {
	return c.conn.Close()
}
