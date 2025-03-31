package user

import (
	"github.com/Masterminds/squirrel"
	"github.com/aridae/goph-keeper/internal/server/database"
)

const (
	uniqueUsernameConstraintName = "users_username_key"
)

const (
	usernameCol     = "username"
	passwordHashCol = "password_hash"
	createdAtCol    = "created_at"
	isDeletedCol    = "is_deleted"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var baseSelectQuery = psql.Select(
	usernameCol,
	passwordHashCol,
).From(database.UsersTable).Where(squirrel.Eq{isDeletedCol: false})
