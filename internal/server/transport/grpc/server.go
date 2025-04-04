package grpc

import (
	"context"
	"fmt"
	"github.com/aridae/goph-keeper/internal/logger"
	gophkeepersecretpb "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/secret"
	gophkeeperuserpb "github.com/aridae/goph-keeper/pkg/pb/goph-keeper/user"
	"google.golang.org/grpc"
	"net"
)

type grpcServer interface {
	grpc.ServiceRegistrar
	Stop()
	Serve(net.Listener) error
}

type Server struct {
	server grpcServer
	port   int
}

func NewServer(
	port int,
	usersAPIServer gophkeeperuserpb.UsersServiceServer,
	secretsAPIServer gophkeepersecretpb.SecretsServiceServer,
	interceptors ...grpc.UnaryServerInterceptor,
) *Server {
	grpcServ := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors...))

	gophkeeperuserpb.RegisterUsersServiceServer(grpcServ, usersAPIServer)
	gophkeepersecretpb.RegisterSecretsServiceServer(grpcServ, secretsAPIServer)

	return &Server{
		server: grpcServ,
		port:   port,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.server.Stop()
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to start tcp listener <port:%d>: %w", s.port, err)
	}
	logger.Infof("start listening on address %v", listener.Addr())

	if err = s.server.Serve(listener); err != nil {
		return fmt.Errorf("failed to start grpc server: %w", err)
	}

	return nil
}
