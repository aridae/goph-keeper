package errormappingmw

import (
	"context"
	"errors"
	grpcerrormapping "github.com/aridae/goph-keeper/internal/common/grpc-error-mapping"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type mockUnaryInvoker struct {
	callError error
}

func (m *mockUnaryInvoker) InvokeFunc() grpc.UnaryInvoker {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return m.callError
	}
}

func TestErrorMapperInterceptor(t *testing.T) {
	testCases := []struct {
		name                string
		invokeError         error
		expectedDomainError error
	}{
		{
			name:                "Grpc Internal Error",
			invokeError:         status.Error(codes.Internal, "internal server error"),
			expectedDomainError: grpcerrormapping.MapGrpcToDomainError(status.Error(codes.Internal, "internal server error")),
		},
		{
			name:                "Grpc Not Found Error",
			invokeError:         status.Error(codes.NotFound, "resource not found"),
			expectedDomainError: grpcerrormapping.MapGrpcToDomainError(status.Error(codes.NotFound, "resource not found")),
		},
		{
			name:                "Custom Error",
			invokeError:         errors.New("custom error"),
			expectedDomainError: grpcerrormapping.MapGrpcToDomainError(errors.New("custom error")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interceptor := ErrorMapperInterceptor()
			ctx := context.Background()
			invoker := &mockUnaryInvoker{callError: tc.invokeError}

			actualError := interceptor(ctx, "method", nil, nil, nil, invoker.InvokeFunc(), nil)

			require.Equal(t, tc.expectedDomainError, actualError)
		})
	}
}
