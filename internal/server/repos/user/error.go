package user

import "errors"

var (
	ErrUsernameUniqueConstraintViolated = errors.New("username unique constraint violation")
)
