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
	em := &Transaction{}

	em.dbCon = GetDbConnection()
	var err error
	em.tx, err = em.dbCon.BeginTxx(ctx, options)
	if err != nil {
		if em.tx != nil {
			_ = em.tx.Rollback()
		}
		return nil, status.Error(codes.Code(500), err.Error())
	}

	return em, nil
}

func (tr *Transaction) PersistNamedCtx(ctx context.Context, query string, entity interface{}) error {
	tx := tr.tx

	_, err := tx.NamedExecContext(ctx, query, entity)

	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}
	return nil
}

func (tr *Transaction) Rollback() error {
	tx := tr.tx

	err := tx.Rollback()
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func (tr *Transaction) Flush() error {
	tx := tr.tx

	err := tx.Commit()
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}
