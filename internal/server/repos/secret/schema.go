package secret

import (
	"github.com/Masterminds/squirrel"
	"github.com/aridae/goph-keeper/internal/server/database"
)

const (
	uniqueOwnerUsernameKeyCombinationConstraintName = "idx_unq_secrets__owner_username__key"
)

const (
	keyCol           = "key"
	dataCol          = "data"
	ownerUsernameCol = "owner_username"
	createdAtCol     = "created_at"
	updatedAtCol     = "updated_at"
	isDeletedCol     = "is_deleted"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var baseSelectQuery = psql.Select(
	keyCol,
	dataCol,
	ownerUsernameCol,
	createdAtCol,
	updatedAtCol,
).From(database.SecretsTable).Where(squirrel.Eq{isDeletedCol: false})
