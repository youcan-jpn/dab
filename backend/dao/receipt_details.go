// This code was written manually.

package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"

	"github.com/youcan-jpn/dab/backend/src/dberror"
)

var ReceiptDetailAllCulumns = []string{
	"receipt_id",
	"shop_id",
	"shop_name",
	"currency_id",
	"currency_name",
	"total_price",
	"purchase_date",
	"product_id",
	"product_name",
	"price",
	"modified_at",
	"created_at",
}

var ReceiptDetailColumnsWOMagic = []string{
	"receipt_id",
	"shop_id",
	"shop_name",
	"currency_id",
	"currency_name",
	"total_price",
	"purchase_date",
	"product_id",
	"product_name",
	"price",
	"modified_at",
	"created_at",
}

type ReceiptDetail struct {
	ReceiptID    int
	ShopID       int
	ShopName     string
	CurrencyID   int
	CurrencyName int
	TotalPrice   float32
	PurchaseDate *time.Time
	ProductID    int
	ProductName  string
	Price        float32
	ModifiedAt   *time.Time
	CreatedAt    *time.Time
}

func (t *ReceiptDetail) Values() []interface{} {
	return []interface{}{
		t.ReceiptID,
		t.ShopID,
		t.ShopName,
		t.CurrencyID,
		t.CurrencyName,
		t.TotalPrice,
		t.PurchaseDate,
		t.ProductID,
		t.ProductName,
		t.Price,
		t.ModifiedAt,
		t.CreatedAt,
	}
}

func (t *ReceiptDetail) SetMap() map[string]interface{} {
	return map[string]interface{}{
		"receipt_id":    t.ReceiptID,
		"shop_id":       t.ShopID,
		"shop_name":     t.ShopName,
		"currency_id":   t.CurrencyID,
		"total_price":   t.TotalPrice,
		"purchase_date": t.PurchaseDate,
		"product_id":    t.ProductID,
		"product_name":  t.ProductName,
		"price":         t.Price,
		"modified_at":   t.ModifiedAt,
		"created_at":    t.CreatedAt,
	}
}

func (t *ReceiptDetail) Ptrs() []interface{} {
	return []interface{}{
		&t.ReceiptID,
		&t.ShopID,
		&t.ShopName,
		&t.CurrencyID,
		&t.CurrencyName,
		&t.TotalPrice,
		&t.PurchaseDate,
		&t.ProductID,
		&t.ProductName,
		&t.Price,
		&t.ModifiedAt,
		&t.CreatedAt,
	}
}

func IterateReceiptDetail(sc interface{ Scan(...interface{}) error }) (ReceiptDetail, error) {
	t := ReceiptDetail{}
	if err := sc.Scan(t.Ptrs()...); err != nil {
		return ReceiptDetail{}, dberror.MapError(err)
	}
	return t, nil
}

func SelectReceiptDetailsByShopIDAndCurrencyIDAndTimeRangeAndPriceRange(ctx context.Context, txn *sql.Tx, shop_id int, currency_id int, min_price float32, max_price float32, since *time.Time, until *time.Time) ([]*ReceiptDetail, error) {
	eq := squirrel.Eq{}
	geq := squirrel.GtOrEq{}
	leq := squirrel.LtOrEq{}

	if shop_id > 0 {
		eq["s.shop_id"] = shop_id
	}
	if currency_id > 0 {
		eq["c.currency_id"] = currency_id
	}
	if min_price > 0 {
		geq["r.total_prcie"] = min_price
	}
	if max_price > 0 {
		leq["r.total_price"] = max_price
	}
	if since != nil {
		geq["r.since"] = since // TODO: この形式で大丈夫かテスト
	}
	if until != nil {
		geq["r.until"] = until // TODO: この形式で大丈夫かテスト
	}

	query, params, err := squirrel.
		Select(
			"r.receipt_id",
			"s.shop_id",
			"s.shop_name",
			"c.currency_id",
			"c.currency_name",
			"r.total_price",
			"r.purchase_date",
			"p.product_id",
			"p.product_name",
			"p.price",
			"r.modified_at",
			"r.created_at",
		).
		From("receipts as r").
		LeftJoin("products as p ON p.receipt_id = r.receipt_id").
		LeftJoin("shops as s ON r.shop_id = s.shop_id").
		LeftJoin("currencies as c ON r.currency_id = c.currency_id").
		Where(eq).
		Where(leq).
		Where(geq).
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
	res := make([]*ReceiptDetail, 0)
	for rows.Next() {
		t, err := IterateReceiptDetail(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}
