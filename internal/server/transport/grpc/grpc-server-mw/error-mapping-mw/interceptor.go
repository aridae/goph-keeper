package errormappingmw

import (
	"context"
	grpcerrormapping "github.com/aridae/goph-keeper/internal/common/grpc-error-mapping"
	"google.golang.org/grpc"
)

func ErrorMapperInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		gRPCError := grpcerrormapping.MapDomainToGrpcError(err)

		return resp, gRPCError
	}
}
