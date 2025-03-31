package secret

import "github.com/aridae/goph-keeper/internal/server/models"

type secretDTO struct {
	OwnerUsername string            `db:"owner_username"`
	Key           string            `db:"key"`
	Data          []byte            `db:"data"`
	Meta          map[string]string `db:"meta"`
}

func mapDTOToDomainSecret(dto secretDTO) models.Secret {
	return models.Secret{
		Accessor: models.SecretAccessor{
			OwnerUsername: dto.OwnerUsername,
			Key:           dto.Key,
		},
		Data: dto.Data,
		Meta: dto.Meta,
	}
}
