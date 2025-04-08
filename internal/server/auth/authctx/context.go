package authctx

import (
	"context"
	"github.com/aridae/goph-keeper/internal/server/models"
)

type contextKey string

const (
	_userCtxKey contextKey = "USER_CONTEXT_KEY"
)

func ContextWithUser(ctx context.Context, user models.User) context.Context {
	return context.WithValue(ctx, _userCtxKey, user)
}

func GetUserFromContext(ctx context.Context) (models.User, bool) {
	user, ok := ctx.Value(_userCtxKey).(models.User)
	if !ok {
		return models.User{}, false
	}

	return user, true
}
