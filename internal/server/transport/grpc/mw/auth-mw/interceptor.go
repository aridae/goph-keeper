package authmw

import (
	"context"
	"github.com/aridae/goph-keeper/internal/server/auth/authctx"
	"github.com/aridae/goph-keeper/internal/server/models"
	"github.com/aridae/goph-keeper/internal/server/pkg/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type jwtService interface {
	ParseToken(ctx context.Context, tokenString string) (jwt.Claims, error)
}

func AuthServerInterceptor(
	jwtService jwtService,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		token, err := extractTokenFromMetaData(ctx, "Bearer")
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		claims, err := jwtService.ParseToken(ctx, token)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}
		user := models.User{Username: claims.Subject}

		ctx = authctx.ContextWithUser(ctx, user)

		return handler(ctx, req)
	}
}

func extractTokenFromMetaData(ctx context.Context, expectedAuthScheme string) (string, error) {
	vals := metadata.ValueFromIncomingContext(ctx, "authorization")
	if len(vals) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with %s", expectedAuthScheme)
	}

	tokenScheme, token, found := strings.Cut(vals[0], " ")
	if !found {
		return "", status.Error(codes.Unauthenticated, "Bad authorization string")
	}

	if !strings.EqualFold(expectedAuthScheme, tokenScheme) {
		return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with %s", expectedAuthScheme)
	}

	return token, nil
}
