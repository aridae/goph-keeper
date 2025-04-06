package jwt

import (
	"context"
	"fmt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Subject string
}

type Service struct {
	secretKeyProvider func(ctx context.Context) []byte
}

func NewService(
	secretKeyProvider func(ctx context.Context) []byte,
) *Service {
	return &Service{secretKeyProvider: secretKeyProvider}
}

func (s *Service) GenerateToken(ctx context.Context, clms Claims) (string, error) {
	claims := jwtv5.MapClaims{
		"sub": clms.Subject,
	}

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)

	signedTokenString, err := token.SignedString(s.secretKeyProvider(ctx))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedTokenString, nil
}

func (s *Service) ParseToken(ctx context.Context, tokenString string) (Claims, error) {
	claims := jwtv5.MapClaims{}

	token, err := jwtv5.ParseWithClaims(tokenString, &claims, func(token *jwtv5.Token) (interface{}, error) { return s.secretKeyProvider(ctx), nil })
	if err != nil {
		return Claims{}, fmt.Errorf("failed to parse token: %w", err)
	}

	subject, err := token.Claims.GetSubject()
	if err != nil {
		return Claims{}, fmt.Errorf("failed to get subject from token claims: %w", err)
	}

	return Claims{Subject: subject}, nil
}
