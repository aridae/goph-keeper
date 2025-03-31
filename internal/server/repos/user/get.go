package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/aridae/goph-keeper/internal/server/models"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetByUsername(ctx context.Context, username string) (*models.UserCredentials, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery.Where(squirrel.Eq{usernameCol: username})

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var dto userDTO
	err = queryable.QueryRow(ctx, sql, args...).Scan(&dto)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to scan row into userDTO: %w", err)
	}

	user := mapDTOToDomainUserCredentials(dto)

	return &user, nil
}
