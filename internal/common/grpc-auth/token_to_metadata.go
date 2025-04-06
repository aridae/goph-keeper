package grpcauth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func PutBearerTokenToMetadata(ctx context.Context, token string) context.Context {
	return putTokenToMetadata(ctx, bearerSchema, token)
}

func putTokenToMetadata(ctx context.Context, authScheme string, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, authorizationHeader, fmt.Sprintf("%s %s", authScheme, token))
}
