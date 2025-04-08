package grpcauth

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractBearerTokenFromMetadata(t *testing.T) {
	tests := []struct {
		name          string
		md            metadata.MD
		expectedToken string
		expectedErr   error
	}{
		{
			name: "Valid Bearer Token",
			md: metadata.Pairs(
				authorizationHeader, "Bearer valid-token",
			),
			expectedToken: "valid-token",
			expectedErr:   nil,
		},
		{
			name:          "Missing Authorization Header",
			md:            metadata.Pairs(),
			expectedToken: "",
			expectedErr:   fmt.Errorf("request unauthenticated with %s", bearerSchema),
		},
		{
			name: "Bad Authorization String",
			md: metadata.Pairs(
				authorizationHeader, "Invalidauthstring",
			),
			expectedToken: "",
			expectedErr:   errors.New("bad authorization string"),
		},
		{
			name: "Wrong Auth Scheme",
			md: metadata.Pairs(
				authorizationHeader, "Basic invalid-token",
			),
			expectedToken: "",
			expectedErr:   fmt.Errorf("request unauthenticated with %s", bearerSchema),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			token, err := ExtractBearerTokenFromMetadata(ctx)

			require.Equal(t, tt.expectedToken, token)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_extractTokenFromMetaData(t *testing.T) {
	tests := []struct {
		name               string
		md                 metadata.MD
		expectedAuthScheme string
		expectedToken      string
		expectedErr        error
	}{
		{
			name: "Valid Basic Token",
			md: metadata.Pairs(
				authorizationHeader, "Basic dXNlcjpwYXNzd29yZA==",
			),
			expectedAuthScheme: "Basic",
			expectedToken:      "dXNlcjpwYXNzd29yZA==",
			expectedErr:        nil,
		},
		{
			name:               "Empty Metadata",
			md:                 metadata.Pairs(),
			expectedAuthScheme: "Basic",
			expectedToken:      "",
			expectedErr:        fmt.Errorf("request unauthenticated with %s", "Basic"),
		},
		{
			name: "Malformed Authorization Header",
			md: metadata.Pairs(
				authorizationHeader, "invalid-auth-string",
			),
			expectedAuthScheme: "Basic",
			expectedToken:      "",
			expectedErr:        errors.New("bad authorization string"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)
			token, err := extractTokenFromMetaData(ctx, tt.expectedAuthScheme)

			require.Equal(t, tt.expectedToken, token)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
