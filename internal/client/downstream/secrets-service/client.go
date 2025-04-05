package secretsservice

import (
	"fmt"
	desc "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	grpc desc.SecretsServiceClient
}

// NewClient создает клиент.
func NewClient(target string) (*Client, error) {
	cc, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial <target:%s>: %w", target, err)
	}

	return &Client{
		conn: cc,
		grpc: desc.NewSecretsServiceClient(cc),
	}, nil
}

// Close закрывает соединение с сервисом.
func (c *Client) Close() error {
	return c.conn.Close()
}
