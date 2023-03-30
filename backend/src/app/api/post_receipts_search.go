package api

import (
	"database/sql"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) PostReceiptsSearch(ctx echo.Context) error {
	var q oapi.PostReceiptsSearchJSONRequestBody
	err := ctx.Bind(&q)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	txn, err := a.dbClient.BeginTx(a.ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	since := objToTime(q.PurchasedSince, a.timezone)
	until := objToTime(q.PurchasedUntil, a.timezone)
	rds, err := service.PostReceiptsSearch(a.ctx, txn, 0, *q.ShopId, *q.CurrencyId, *q.CategoryId, *q.MinPrice, *q.MaxPrice, since, until)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if len(rds) == 0 {
		return ctx.JSON(http.StatusOK, oapi.Receipts{})
	}

	sort.Slice(rds, func(i, j int) bool {
		return rds[i].ReceiptID*10000+rds[i].ProductID < rds[j].ReceiptID*10000+rds[j].ProductID
	}) // 各receiptにproductは9999個以下と仮定
	var r oapi.Receipt
	rid := -1
	rs := []oapi.Receipt{}
	for _, rd := range rds {
		if rid != rd.ReceiptID {
			rid = rd.ReceiptID
			r = oapi.Receipt{
				ReceiptId:    &rd.ReceiptID,
				ShopId:       &rd.ShopID,
				ShopName:     &rd.ShopName,
				CurrencyId:   &rd.CurrencyID,
				CurrencyName: &rd.CurrencyName,
				CategoryId:   &rd.CategoryID,
				CategoryName: &rd.CategoryName,
				TotalPrice:   &rd.TotalPrice,
				PurchaseDate: timeToObj(rd.PurchaseDate),
				Products:     &[]oapi.Product{},
			}
			rs = append(rs, r)
		}
		p := oapi.Product{
			ProductId:   &rd.ProductID,
			ProductName: &rd.ProductName,
			Price:       &rd.Price,
		}
		*r.Products = append(*r.Products, p)
	}

	return ctx.JSON(http.StatusOK, oapi.Receipts{Receipts: &rs})
}
