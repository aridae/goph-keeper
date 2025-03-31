package secret

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

func (r *Repository) CreateSecret(ctx context.Context, secret models.Secret, now time.Time) error {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := psql.Insert(database.SecretsTable).
		Columns(
			keyCol,
			ownerUsernameCol,
			dataCol,
			createdAtCol,
			updatedAtCol,
		).
		Values(
			secret.Accessor.Key,
			secret.Accessor.OwnerUsername,
			secret.Data,
			now,
			now,
		).
		Suffix("returning id")

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	if _, err = queryable.Exec(ctx, sql, args...); err != nil {
		if pgerr := new(pgconn.PgError); errors.As(err, &pgerr) && isOwnerUsernameKeyCombinationConstraintViolated(pgerr) {
			return ErrOwnerUsernameKeyCombinationConstraintViolated
		}

		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func isOwnerUsernameKeyCombinationConstraintViolated(pgerr *pgconn.PgError) bool {
	return pgerr.Code == pgerrcode.UniqueViolation && pgerr.ConstraintName == uniqueOwnerUsernameKeyCombinationConstraintName
}
