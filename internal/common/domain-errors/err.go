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
	Msg  string
	Code int64
}

func (de DomainError) Error() string {
	return de.Msg
}

func ErrUnauthorized() error {
	return DomainError{
		Msg:  "Unauthorized",
		Code: UnauthorizedErrorCode,
	}
}

func ErrInvalidUserCredentials() error {
	return DomainError{
		Msg:  "Invalid credentials",
		Code: InvalidArgumentErrorCode,
	}
}

func ErrUsernameAlreadyTaken(username string) error {
	return DomainError{
		Msg:  fmt.Sprintf("Username %s already taken", username),
		Code: FailedPreconditionErrorCode,
	}
}

func ErrSecretNotFound(key string) error {
	return DomainError{
		Msg:  fmt.Sprintf("Secret %s not found", key),
		Code: NotFoundErrorCode,
	}
}

func ErrSecretAlreadyExists(key string) error {
	return DomainError{
		Msg:  fmt.Sprintf("Secret %s already exists", key),
		Code: FailedPreconditionErrorCode,
	}
}
