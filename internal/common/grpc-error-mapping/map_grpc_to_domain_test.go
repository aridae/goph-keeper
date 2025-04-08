package grpcerrormapping

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	domainerrors "github.com/aridae/goph-keeper/internal/common/domain-errors"
)

func TestMapGrpcToDomainError(t *testing.T) {
	tests := []struct {
		name           string
		inputError     error
		expectedDomain domainerrors.DomainError
	}{
		{
			name:       "Invalid Argument Error",
			inputError: status.Error(codes.InvalidArgument, "Invalid argument"),
			expectedDomain: domainerrors.DomainError{
				Msg:  "Invalid argument",
				Code: domainerrors.InvalidArgumentErrorCode,
			},
		},
		{
			name:       "Unauthenticated Error",
			inputError: status.Error(codes.Unauthenticated, "Unauthorized access"),
			expectedDomain: domainerrors.DomainError{
				Msg:  "Unauthorized access",
				Code: domainerrors.UnauthorizedErrorCode,
			},
		},
		{
			name:       "Failed Precondition Error",
			inputError: status.Error(codes.FailedPrecondition, "Precondition failed"),
			expectedDomain: domainerrors.DomainError{
				Msg:  "Precondition failed",
				Code: domainerrors.FailedPreconditionErrorCode,
			},
		},
		{
			name:       "Not Found Error",
			inputError: status.Error(codes.NotFound, "Resource not found"),
			expectedDomain: domainerrors.DomainError{
				Msg:  "Resource not found",
				Code: domainerrors.NotFoundErrorCode,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapGrpcToDomainError(tt.inputError)

			require.Equal(t, tt.expectedDomain, result)
		})
	}
}
