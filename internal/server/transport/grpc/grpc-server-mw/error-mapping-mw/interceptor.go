package errormappingmw

import (
	"context"
	"errors"
	"github.com/aridae/goph-keeper/internal/logger"
	domainerrors "github.com/aridae/goph-keeper/internal/server/models/domain-errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type logFn func(format string, args ...interface{})

// mapError разворачивает ошибку и присваивает ей gRPC код и уровень логирования.
func mapError(err error) (logFn, error) {
	if businessError := new(domainerrors.DomainError); errors.As(err, businessError) {
		switch businessError.Code {
		case domainerrors.InvalidArgumentErrorCode:
			return logger.Infof, status.Error(codes.InvalidArgument, businessError.Error())
		case domainerrors.NotFoundErrorCode:
			return logger.Infof, status.Error(codes.NotFound, businessError.Error())
		case domainerrors.FailedPreconditionErrorCode:
			return logger.Infof, status.Error(codes.FailedPrecondition, businessError.Error())
		case domainerrors.UnauthorizedErrorCode:
			return logger.Infof, status.Error(codes.Unauthenticated, businessError.Error())
		default:
			logger.Warnf("servermw.mapError: unknown tms-lib <errorCode:%d>", businessError.Code)
		}
	}
	if errors.Is(err, context.Canceled) {
		// keeping it as debug, because we don't want to see it most of the time
		return logger.Debugf, status.Error(codes.Canceled, "context cancelled by caller")
	}

	return logger.Errorf, status.Error(codes.Internal, err.Error())
}

// ErrorHandlerMw grpc middleware, обрабатывающее ошибки и производящее их логирование.
func ErrorHandlerMw() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		var fullMethodName string
		if info != nil {
			fullMethodName = info.FullMethod
		}

		logFunc, gRPCError := mapError(err)
		st, _ := status.FromError(gRPCError)

		logFunc("'%s' упал с кодом '%s': %v", fullMethodName, st.Code().String(), err)

		return resp, gRPCError
	}
}
