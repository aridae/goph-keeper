package grpcauth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

// PutBearerTokenToMetadata добавляет токен Bearer в метаданные контекста gRPC.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// token string - токен Bearer для добавления.
//
// Возвращаемые значения:
// context.Context - обновлённый контекст с добавленным токеном.
func PutBearerTokenToMetadata(ctx context.Context, token string) context.Context {
	return putTokenToMetadata(ctx, bearerSchema, token)
}

// putTokenToMetadata добавляет токен с заданной схемой авторизации в метаданные контекста gRPC.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// authScheme string - схема авторизации (например, "Bearer").
// token string - токен для добавления.
//
// Возвращаемые значения:
// context.Context - обновлённый контекст с добавленным токеном.
func putTokenToMetadata(ctx context.Context, authScheme string, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, authorizationHeader, fmt.Sprintf("%s %s", authScheme, token))
}
