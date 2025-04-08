package authmw

import (
	"context"
	"github.com/aridae/goph-keeper/internal/client/downstream/grpc-client-mw/auth-mw/_mock"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestAuthInterceptor(t *testing.T) {
	testCases := []struct {
		name  string
		token *string
	}{
		{
			name: "No Token Available",
		},
		{
			name:  "Token Available",
			token: &[]string{"TOKEN"}[0],
		},
	}

	invoker := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return nil
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			sessionStorageMock := _mock.NewMocksessionStorage(ctrl)

			interceptor := AuthInterceptor(sessionStorageMock)

			sessionStorageMock.EXPECT().GetToken(gomock.Any()).Return(tc.token)

			err := interceptor(ctx, "method", nil, nil, nil, invoker)

			require.NoError(t, err)
		})
	}
}
