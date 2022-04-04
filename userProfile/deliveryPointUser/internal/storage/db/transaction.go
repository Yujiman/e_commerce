package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Transaction struct {
	dbCon *sqlx.DB
	tx    *sqlx.Tx
}

func NewTransaction(ctx context.Context, options *sql.TxOptions) (*Transaction, error) {
	tr := &Transaction{}

	tr.dbCon = GetDbConnection()
	var err error

	tr.tx, err = tr.dbCon.BeginTxx(ctx, options)
	if err != nil {
		if tr.tx != nil {
			_ = tr.tx.Rollback()
		}
		return nil, status.Error(codes.Code(500), err.Error())
	}

	return tr, nil
}

func (tr *Transaction) PersistNamedCtx(ctx context.Context, query string, entity interface{}) error {
	_, err := tr.tx.NamedExecContext(ctx, query, entity)
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (tr *Transaction) PersistExecContext(ctx context.Context, query string) error {
	_, err := tr.tx.ExecContext(ctx, query)
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (tr *Transaction) Rollback() error {
	if err := tr.tx.Rollback(); err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (tr *Transaction) Flush() error {
	if err := tr.tx.Commit(); err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}
