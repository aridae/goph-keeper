package jwt

import (
	"context"
	"fmt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// Claims аутентификационные данные JWT-токена.
type Claims struct {
	Subject string
}

// Service предоставляет функциональность для генерации и парсинга JWT-токенов.
type Service struct {
	secretKeyProvider func(ctx context.Context) []byte
}

// NewService создаёт новый экземпляр сервиса для работы с JWT-токенами.
//
// Параметры:
// secretKeyProvider func(ctx context.Context) []byte - функция, предоставляющая секретный ключ.
//
// Возвращаемые значения:
// *Service - новый экземпляр сервиса.
func NewService(
	secretKeyProvider func(ctx context.Context) []byte,
) *Service {
	return &Service{secretKeyProvider: secretKeyProvider}
}

// GenerateToken генерирует JWT-токен с указанными аутентификационными данными.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// clms Claims - аутентификационные данные токена.
//
// Возвращаемые значения:
// string - строка с подписанным JWT-токеном.
// error - ошибка, если токен не удалось подписать.
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

// ParseToken парсит JWT-токен и возвращает аутентификационные данные.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// tokenString string - строка с JWT-токеном.
//
// Возвращаемые значения:
// Claims - аутентификационные данные токена.
// error - ошибка, если токен не удалось распарсить.
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
