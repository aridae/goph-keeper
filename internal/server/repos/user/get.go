package user

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/aridae/goph-keeper/internal/server/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *Repository) GetByUsername(ctx context.Context, username string) (*models.UserCredentials, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery.Where(squirrel.Eq{usernameCol: username})

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := queryable.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	var dtos []userDTO
	err = pgxscan.ScanAll(&dtos, rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan row into userDTO: %w", err)
	}

	if len(dtos) == 0 {
		return nil, nil
	}
	dto := dtos[0]

	user := mapDTOToDomainUserCredentials(dto)

	return &user, nil
}
