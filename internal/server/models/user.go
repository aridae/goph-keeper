package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentials struct {
	Username     string
	PasswordHash []byte
}

func NewUserCredentials(uname string, pass string) (UserCredentials, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return UserCredentials{}, fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}

	return UserCredentials{
		Username:     uname,
		PasswordHash: pwdHash,
	}, nil
}

func (c UserCredentials) Equal(uname string, pass string) bool {
	if c.Username != uname {
		return false
	}

	if err := bcrypt.CompareHashAndPassword(c.PasswordHash, []byte(pass)); err != nil {
		return false
	}

	return true
}

type User struct {
	Username string
}
