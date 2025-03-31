package secret

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/aridae/goph-keeper/internal/server/models"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetByAccessor(ctx context.Context, accessor models.SecretAccessor) (*models.Secret, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery.
		Where(squirrel.Eq{ownerUsernameCol: accessor.OwnerUsername}).
		Where(squirrel.Eq{keyCol: accessor.Key})

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var dto secretDTO
	err = queryable.QueryRow(ctx, sql, args...).Scan(&dto)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to scan row into userDTO: %w", err)
	}

	user := mapDTOToDomainSecret(dto)

	return &user, nil
}
