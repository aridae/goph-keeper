package user

import (
	"github.com/aridae/goph-keeper/internal/server/models"
)

type userDTO struct {
	Username     string `db:"username"`
	PasswordHash []byte `db:"password_hash"`
}

func mapDTOToDomainUserCredentials(dto userDTO) models.UserCredentials {
	return models.UserCredentials{
		Username:     dto.Username,
		PasswordHash: dto.PasswordHash,
	}
}
