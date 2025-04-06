package jwt

import (
	"context"
	"testing"

	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	secretKey := []byte("super-secret-key")

	service := NewService(func(ctx context.Context) []byte {
		return secretKey
	})

	clms := Claims{
		Subject: "test-user",
	}

	ctx := context.Background()

	tokenStr, err := service.GenerateToken(ctx, clms)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)

	parsedToken, err := jwtv5.Parse(tokenStr, func(token *jwtv5.Token) (interface{}, error) {
		return secretKey, nil
	})
	require.NoError(t, err)
	require.NotNil(t, parsedToken)

	claims, ok := parsedToken.Claims.(jwtv5.MapClaims)
	require.True(t, ok)
	require.Equal(t, "test-user", claims["sub"])
}

func TestParseToken(t *testing.T) {
	secretKey := []byte("super-secret-key")

	service := NewService(func(ctx context.Context) []byte {
		return secretKey
	})

	clms := Claims{
		Subject: "test-user",
	}

	ctx := context.Background()

	tokenStr, err := service.GenerateToken(ctx, clms)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)

	parsedClms, err := service.ParseToken(ctx, tokenStr)
	require.NoError(t, err)
	require.Equal(t, clms, parsedClms)
}

func TestParseToken_InvalidSignature(t *testing.T) {
	service := NewService(func(ctx context.Context) []byte {
		return []byte("wrong-secret-key")
	})

	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0LXVzZXIifQ.wrong_signature"

	ctx := context.Background()

	_, err := service.ParseToken(ctx, tokenStr)
	require.Error(t, err)
	require.Contains(t, err.Error(), "signature is invalid")
}
