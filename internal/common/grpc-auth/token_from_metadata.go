package grpcauth

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authorizationHeader = "authorization"
	bearerSchema        = "Bearer"
)

func ExtractBearerTokenFromMetadata(ctx context.Context) (string, error) {
	return extractTokenFromMetaData(ctx, bearerSchema)
}

func extractTokenFromMetaData(ctx context.Context, expectedAuthScheme string) (string, error) {
	vals := metadata.ValueFromIncomingContext(ctx, authorizationHeader)
	if len(vals) == 0 {
		return "", fmt.Errorf("request unauthenticated with %s", expectedAuthScheme)
	}

	tokenScheme, token, found := strings.Cut(vals[0], " ")
	if !found {
		return "", errors.New("bad authorization string")
	}

	if !strings.EqualFold(expectedAuthScheme, tokenScheme) {
		return "", fmt.Errorf("request unauthenticated with %s", expectedAuthScheme)
	}

	return token, nil
}
