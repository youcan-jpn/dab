// This code was written manually.

package dao

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"

	"github.com/youcan-jpn/dab/backend/gen/daocore"
	"github.com/youcan-jpn/dab/backend/src/dberror"
)

func InsertOneCategoryReturningResult(ctx context.Context, txn *sql.Tx, record *daocore.Category) (sql.Result, error) {
	query, params, err := squirrel.
		Insert(daocore.CategoryTableName).
		Columns(daocore.CategoryColumnsWOMagics...).
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
