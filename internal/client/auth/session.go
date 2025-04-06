package auth

import (
	"context"
	"sync"
)

const (
	accessTokenStoreKey = "access-token-store-key"
)

// Session представляет собой объект для управления токеном сессии.
type Session struct {
	store sync.Map
}

// NewSession создает новый экземпляр Session.
//
// Возвращаемое значение:
// *Session - новый экземпляр Session.
func NewSession() *Session {
	return &Session{}
}

// StoreToken сохраняет токен в хранилище сессии.
//
// Параметры:
// ctx context.Context - контекст выполнения.
// token string - токен для сохранения.
//
// Возвращаемые значения:
// error - ошибка, если она возникла при сохранении токена.
func (s *Session) StoreToken(_ context.Context, token string) error {
	s.store.Store(accessTokenStoreKey, token)

	return nil
}

// GetToken извлекает токен из хранилища сессии.
//
// Параметры:
// ctx context.Context - контекст выполнения.
//
// Возвращаемые значения:
// *string - токен, если он был найден, иначе nil.
func (s *Session) GetToken(_ context.Context) *string {
	tokenVal, isAuthorized := s.store.Load(accessTokenStoreKey)
	if !isAuthorized {
		return nil
	}

	tokenString, ok := tokenVal.(string)
	if !ok {
		return nil
	}

	return &tokenString
}
