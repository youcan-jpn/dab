// This code was written manually.

package dao

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"

	"github.com/youcan-jpn/dab/backend/gen/daocore"
	"github.com/youcan-jpn/dab/backend/src/dberror"
)

func SelectAllCurrencies(ctx context.Context, txn *sql.Tx) ([]*daocore.Currency, error) {
	query, params, err := squirrel.
		Select(daocore.CurrencyAllColumns...).
		From(daocore.CurrencyTableName).
		ToSql()
	if err != nil {
		return nil, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return nil, dberror.MapError(err)
	}
	rows, err := stmt.QueryContext(ctx, params...)
	if err != nil {
		return nil, dberror.MapError(err)
	}
	res := make([]*daocore.Currency, 0)
	for rows.Next() {
		t, err := daocore.IterateCurrency(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}
