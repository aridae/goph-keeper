package auth

import (
	"context"
	"sync"
)

const (
	accessTokenStoreKey = "access-token-store-key"
)

type Session struct {
	store sync.Map
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) StoreToken(_ context.Context, token string) error {
	s.store.Store(accessTokenStoreKey, token)

	return nil
}

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
