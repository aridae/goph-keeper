package authmw

import (
	"context"
	grpcauth "github.com/aridae/goph-keeper/internal/common/grpc-auth"
	"github.com/aridae/goph-keeper/internal/server/auth/authctx"
	"github.com/aridae/goph-keeper/internal/server/models"
	"github.com/aridae/goph-keeper/internal/server/pkg/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"slices"
)

type jwtService interface {
	ParseToken(ctx context.Context, tokenString string) (jwt.Claims, error)
}

func AuthServerInterceptor(
	jwtService jwtService,
	whitelist []string,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if slices.Contains(whitelist, info.FullMethod) {
			return handler(ctx, req)
		}

		token, err := grpcauth.ExtractBearerTokenFromMetadata(ctx)
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
