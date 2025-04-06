package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	session := NewSession()
	assert.NotNil(t, session)
	assert.IsType(t, &Session{}, session)
}

func TestStoreToken(t *testing.T) {
	ctx := context.Background()
	session := NewSession()

	err := session.StoreToken(ctx, "test-token")
	assert.NoError(t, err)

	token := session.GetToken(ctx)
	assert.Equal(t, "test-token", *token)
}

func TestGetToken_NoTokenStored(t *testing.T) {
	ctx := context.Background()
	session := NewSession()

	token := session.GetToken(ctx)
	assert.Nil(t, token)
}

func TestGetToken_InvalidTokenType(t *testing.T) {
	ctx := context.Background()
	session := NewSession()

	session.store.Store(accessTokenStoreKey, 12345)

	token := session.GetToken(ctx)
	assert.Nil(t, token)
}
