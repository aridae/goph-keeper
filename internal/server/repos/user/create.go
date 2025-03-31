package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/aridae/goph-keeper/internal/server/database"
	"github.com/aridae/goph-keeper/internal/server/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

func (r *Repository) CreateUser(ctx context.Context, user models.UserCredentials, now time.Time) error {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := psql.Insert(database.UsersTable).
		Columns(
			usernameCol,
			passwordHashCol,
			createdAtCol,
		).
		Values(
			user.Username,
			user.PasswordHash,
			now,
		).
		Suffix("returning id")

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	if _, err = queryable.Exec(ctx, sql, args...); err != nil {
		if pgerr := new(pgconn.PgError); errors.As(err, &pgerr) && isUsernameUniqueConstraintViolated(pgerr) {
			return ErrUsernameUniqueConstraintViolated
		}

		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func isUsernameUniqueConstraintViolated(pgerr *pgconn.PgError) bool {
	return pgerr.Code == pgerrcode.UniqueViolation && pgerr.ConstraintName == uniqueUsernameConstraintName
}
