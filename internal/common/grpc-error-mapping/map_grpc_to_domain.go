package grpcerrormapping

import (
	domainerrors "github.com/aridae/goph-keeper/internal/common/domain-errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapGrpcToDomainError(err error) error {
	grpcStatus, ok := status.FromError(err)
	if !ok {
		return err
	}

	switch grpcStatus.Code() {
	case codes.InvalidArgument:
		return domainerrors.DomainError{
			Msg:  grpcStatus.Message(),
			Code: domainerrors.InvalidArgumentErrorCode,
		}
	case codes.Unauthenticated:
		return domainerrors.DomainError{
			Msg:  grpcStatus.Message(),
			Code: domainerrors.UnauthorizedErrorCode,
		}
	case codes.FailedPrecondition:
		return domainerrors.DomainError{
			Msg:  grpcStatus.Message(),
			Code: domainerrors.FailedPreconditionErrorCode,
		}
	case codes.NotFound:
		return domainerrors.DomainError{
			Msg:  grpcStatus.Message(),
			Code: domainerrors.NotFoundErrorCode,
		}
	default:
		return err
	}
}
