package user

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type sqlQueryable interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type sqlTransactionManager interface {
	DefaultTrOrDB(ctx context.Context, db trmpgx.Tr) trmpgx.Tr
}

type Repository struct {
	db       sqlQueryable
	txGetter sqlTransactionManager
}

func NewRepository(
	db sqlQueryable,
	txGetter sqlTransactionManager,
) *Repository {
	return &Repository{db: db, txGetter: txGetter}
}
