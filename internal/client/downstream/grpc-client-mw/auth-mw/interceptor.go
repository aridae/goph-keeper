package authmw

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type sessionStorage interface {
	GetToken(ctx context.Context) *string
}

func AuthClientInterceptor(
	sessionStorage sessionStorage,
) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		token := sessionStorage.GetToken(ctx)
		if token != nil {
			ctx = putTokenToMetaData(ctx, "Bearer", *token)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func putTokenToMetaData(ctx context.Context, authScheme string, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("%s %s", authScheme, token))
}
