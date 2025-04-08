package errormappingmw

import (
	"context"
	grpcerrormapping "github.com/aridae/goph-keeper/internal/common/grpc-error-mapping"
	"google.golang.org/grpc"
)

// ErrorMapperInterceptor создаёт gRPC UnaryClientInterceptor, который преобразует ошибки gRPC в ошибки доменного уровня.
//
// Возвращаемые значения:
// grpc.UnaryClientInterceptor - интерсептор для клиентских вызовов gRPC.
func ErrorMapperInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)

		domainErr := grpcerrormapping.MapGrpcToDomainError(err)

		return domainErr
	}
}
