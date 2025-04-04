package domainerrors

import (
	"fmt"
)

const (
	UnauthorizedErrorCode = iota + 1
	FailedPreconditionErrorCode
	InvalidArgumentErrorCode
	NotFoundErrorCode
)

type DomainError struct {
	msg  string
	Code int64
}

func (de DomainError) Error() string {
	return de.msg
}

func ErrUnauthorized() error {
	return DomainError{
		msg:  "Unauthorized",
		Code: UnauthorizedErrorCode,
	}
}

func ErrInvalidUserCredentials() error {
	return DomainError{
		msg:  "Invalid credentials",
		Code: InvalidArgumentErrorCode,
	}
}

func ErrUsernameAlreadyTaken(username string) error {
	return DomainError{
		msg:  fmt.Sprintf("Username %s already taken", username),
		Code: FailedPreconditionErrorCode,
	}
}

func ErrSecretNotFound(key string) error {
	return DomainError{
		msg:  fmt.Sprintf("Secret %s not found", key),
		Code: NotFoundErrorCode,
	}
}

func ErrSecretAlreadyExists(key string) error {
	return DomainError{
		msg:  fmt.Sprintf("Secret %s already exists", key),
		Code: FailedPreconditionErrorCode,
	}
}
