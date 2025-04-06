package grpcauth

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestPutBearerTokenToMetadata(t *testing.T) {
	tests := []struct {
		name   string
		ctx    context.Context
		token  string
		wantMD metadata.MD
	}{
		{
			name:   "Add Bearer Token",
			ctx:    context.Background(),
			token:  "my-bearer-token",
			wantMD: metadata.Pairs(authorizationHeader, "Bearer my-bearer-token"),
		},
		{
			name:   "Empty Context",
			ctx:    context.Background(),
			token:  "",
			wantMD: metadata.Pairs(authorizationHeader, "Bearer "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtx := PutBearerTokenToMetadata(tt.ctx, tt.token)
			md, ok := metadata.FromOutgoingContext(gotCtx)
			require.True(t, ok)
			require.Equal(t, tt.wantMD, md)
		})
	}
}

func Test_putTokenToMetadata(t *testing.T) {
	tests := []struct {
		name       string
		ctx        context.Context
		authScheme string
		token      string
		wantMD     metadata.MD
	}{
		{
			name:       "Add Basic Token",
			ctx:        context.Background(),
			authScheme: "Basic",
			token:      "my-basic-token",
			wantMD:     metadata.Pairs(authorizationHeader, "Basic my-basic-token"),
		},
		{
			name:       "Empty Token",
			ctx:        context.Background(),
			authScheme: "Basic",
			token:      "",
			wantMD:     metadata.Pairs(authorizationHeader, "Basic "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtx := putTokenToMetadata(tt.ctx, tt.authScheme, tt.token)
			md, ok := metadata.FromOutgoingContext(gotCtx)
			require.True(t, ok)
			require.Equal(t, tt.wantMD, md)
		})
	}
}
