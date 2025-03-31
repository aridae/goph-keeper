package domainerrors

import (
	"fmt"
)

const (
	UnauthorizedErrorCode = iota + 1
	UsernameAlreadyTakenErrorCode
	InvalidUserCredentialsErrorCode
	SecretNotFoundErrorCode
	SecretAlreadyExists
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
		Code: InvalidUserCredentialsErrorCode,
	}
}

func ErrUsernameAlreadyTaken(username string) error {
	return DomainError{
		msg:  fmt.Sprintf("Username %s already taken", username),
		Code: UsernameAlreadyTakenErrorCode,
	}
}

func ErrSecretNotFound(key string) error {
	return DomainError{
		msg:  fmt.Sprintf("Secret %s not found", key),
		Code: SecretNotFoundErrorCode,
	}
}

func ErrSecretAlreadyExists(key string) error {
	return DomainError{
		msg:  fmt.Sprintf("Secret %s already exists", key),
		Code: SecretAlreadyExists,
	}
}
