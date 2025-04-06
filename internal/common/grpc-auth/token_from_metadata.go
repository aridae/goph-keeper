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

// ExtractBearerTokenFromMetadata извлекает токен Bearer из метаданных контекста gRPC.
//
// Параметры:
// ctx context.Context - контекст выполнения.
//
// Возвращаемые значения:
// string - токен Bearer, если он был найден.
// error - ошибка, если токен не был найден или произошла другая ошибка.
func ExtractBearerTokenFromMetadata(ctx context.Context) (string, error) {
	return extractTokenFromMetaData(ctx, bearerSchema)
}

// extractTokenFromMetaData извлекает токен из метаданных контекста gRPC, проверяя схему авторизации.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// expectedAuthScheme string - ожидаемая схема авторизации (например, "Bearer").
//
// Возвращаемые значения:
// string - токен, если он был найден.
// error - ошибка, если токен не был найден или произошла другая ошибка.
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
