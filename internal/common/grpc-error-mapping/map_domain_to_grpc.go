package grpcerrormapping

import (
	"context"
	"errors"
	domainerrors "github.com/aridae/goph-keeper/internal/common/domain-errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MapDomainToGrpcError преобразует ошибки доменного уровня в ошибки gRPC.
//
// Параметры:
// err error - ошибка доменного уровня.
//
// Возвращаемые значения:
// error - соответствующая ошибка gRPC.
func MapDomainToGrpcError(err error) error {
	if domainError := new(domainerrors.DomainError); errors.As(err, domainError) {
		switch domainError.Code {
		case domainerrors.InvalidArgumentErrorCode:
			return status.Error(codes.InvalidArgument, domainError.Error())
		case domainerrors.NotFoundErrorCode:
			return status.Error(codes.NotFound, domainError.Error())
		case domainerrors.FailedPreconditionErrorCode:
			return status.Error(codes.FailedPrecondition, domainError.Error())
		case domainerrors.UnauthorizedErrorCode:
			return status.Error(codes.Unauthenticated, domainError.Error())
		}
	}
	if errors.Is(err, context.Canceled) {
		return status.Error(codes.Canceled, "context cancelled by caller")
	}

	return status.Error(codes.Internal, err.Error())
}
