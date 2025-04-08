package authmw

import (
	"context"
	grpcauth "github.com/aridae/goph-keeper/internal/common/grpc-auth"
	"google.golang.org/grpc"
)

type sessionStorage interface {
	GetToken(ctx context.Context) *string
}

// AuthInterceptor создаёт gRPC UnaryClientInterceptor, который добавляет токен Bearer в контекст вызова gRPC.
//
// Параметры:
// sessionStorage sessionStorage - интерфейс для получения токена из хранилища сессий.
//
// Возвращаемые значения:
// grpc.UnaryClientInterceptor - интерсептор для клиентских вызовов gRPC.
func AuthInterceptor(
	sessionStorage sessionStorage,
) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		token := sessionStorage.GetToken(ctx)
		if token != nil {
			ctx = grpcauth.PutBearerTokenToMetadata(ctx, *token)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
