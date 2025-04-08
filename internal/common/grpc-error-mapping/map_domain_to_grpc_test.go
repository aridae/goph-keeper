package grpcerrormapping

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	domainerrors "github.com/aridae/goph-keeper/internal/common/domain-errors"
)

func TestMapDomainToGrpcError(t *testing.T) {
	tests := []struct {
		name           string
		inputError     error
		expectedStatus codes.Code
		expectedMsg    string
	}{
		{
			name: "Invalid Argument Error",
			inputError: domainerrors.DomainError{
				Msg:  "Invalid argument",
				Code: domainerrors.InvalidArgumentErrorCode,
			},
			expectedStatus: codes.InvalidArgument,
			expectedMsg:    "Invalid argument",
		},
		{
			name: "Not Found Error",
			inputError: domainerrors.DomainError{
				Msg:  "Resource not found",
				Code: domainerrors.NotFoundErrorCode,
			},
			expectedStatus: codes.NotFound,
			expectedMsg:    "Resource not found",
		},
		{
			name: "Failed Precondition Error",
			inputError: domainerrors.DomainError{
				Msg:  "Precondition failed",
				Code: domainerrors.FailedPreconditionErrorCode,
			},
			expectedStatus: codes.FailedPrecondition,
			expectedMsg:    "Precondition failed",
		},
		{
			name: "Unauthorized Error",
			inputError: domainerrors.DomainError{
				Msg:  "Unauthorized access",
				Code: domainerrors.UnauthorizedErrorCode,
			},
			expectedStatus: codes.Unauthenticated,
			expectedMsg:    "Unauthorized access",
		},
		{
			name:           "Canceled Context",
			inputError:     context.Canceled,
			expectedStatus: codes.Canceled,
			expectedMsg:    "context cancelled by caller",
		},
		{
			name: "Unknown Domain Error",
			inputError: domainerrors.DomainError{
				Msg:  "Unknown domain error",
				Code: 9999,
			},
			expectedStatus: codes.Internal,
			expectedMsg:    "Unknown domain error",
		},
		{
			name:           "Internal Error",
			inputError:     errors.New("Some internal error"),
			expectedStatus: codes.Internal,
			expectedMsg:    "Some internal error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapDomainToGrpcError(tt.inputError)
			st, ok := status.FromError(result)
			require.True(t, ok)
			require.Equal(t, tt.expectedStatus, st.Code())
			require.Contains(t, st.Message(), tt.expectedMsg)
		})
	}
}
