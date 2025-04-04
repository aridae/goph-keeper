package secret

import (
	"github.com/aridae/goph-keeper/internal/server/models"
	"time"
)

type secretDTO struct {
	OwnerUsername string            `db:"owner_username"`
	Key           string            `db:"key"`
	Data          []byte            `db:"data"`
	Meta          map[string]string `db:"meta"`
	CreatedAt     time.Time         `db:"created_at"`
	UpdatedAt     time.Time         `db:"updated_at"`
}

func mapDTOToDomainSecret(dto secretDTO) models.Secret {
	return models.Secret{
		Accessor: models.SecretAccessor{
			OwnerUsername: dto.OwnerUsername,
			Key:           dto.Key,
		},
		Data:      dto.Data,
		Meta:      dto.Meta,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
