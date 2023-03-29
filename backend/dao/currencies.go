// This code was written manually.

package dao

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"

	"github.com/youcan-jpn/dab/backend/gen/daocore"
	"github.com/youcan-jpn/dab/backend/src/dberror"
)

func InsertOneCurrencyReturningResult(ctx context.Context, txn *sql.Tx, record *daocore.Currency) (sql.Result, error) {
	query, params, err := squirrel.
		Insert(daocore.CurrencyTableName).
		Columns(daocore.CurrencyColumnsWOMagics...).
		Values(record.Values()...).
		ToSql()
	if err != nil {
		return nil, err
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return nil, dberror.MapError(err)
	}
	r, err := stmt.Exec(params...)
	if err != nil {
		return nil, dberror.MapError(err)
	}
	return r, nil
}
