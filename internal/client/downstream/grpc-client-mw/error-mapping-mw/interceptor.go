package errormappingmw

import (
	"context"
	grpcerrormapping "github.com/aridae/goph-keeper/internal/common/grpc-error-mapping"
	"google.golang.org/grpc"
)

func ErrorMapperInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)

		domainErr := grpcerrormapping.MapGrpcToDomainError(err)

		return domainErr
	}
}
