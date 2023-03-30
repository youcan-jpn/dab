package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) PatchReceiptsReceiptId(ctx echo.Context, receiptId oapi.ReceiptId) error {
	var rct oapi.PatchReceiptsReceiptIdJSONRequestBody
	err := ctx.Bind(&rct)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	txn, err := a.dbClient.BeginTx(a.ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	defer txn.Rollback()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	tp := 0.0
	ps := []*service.Product{}
	for _, p := range *rct.Products {
		np := service.Product{
			ProductID:   *p.ProductId,
			ProductName: *p.ProductName,
			Price:       *p.Price,
		}
		ps = append(ps, &np)
		tp += float64(*p.Price)
	}

	pd := objToTime(rct.PurchaseDate, a.timezone)
	err = service.PatchReceipt(a.ctx, txn, service.Receipt{
		ReceiptID:    receiptId,
		ShopID:       *rct.ShopId,
		CurrencyID:   *rct.CurrencyId,
		CategoryID:   *rct.CategoryId,
		TotalPrice:   float32(tp),
		PurchaseDate: pd,
		Products:     ps,
	})
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res, err := service.PostReceiptsSearch(a.ctx, txn, receiptId, 0, 0, 0, 0, 0, nil, nil)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if len(res) < 1 {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("receipt_id (%v) does not exist", receiptId))
	}

	pros := oapi.Products{}
	for _, rp := range res {
		p := oapi.Product{
			ProductId:   &rp.ProductID,
			ProductName: &rp.ProductName,
			Price:       &rp.Price,
		}
		pros = append(pros, p)
	}
	if err = txn.Commit(); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, oapi.Receipt{
		ReceiptId:    &res[0].ReceiptID,
		ShopId:       &res[0].ShopID,
		ShopName:     &res[0].ShopName,
		CurrencyId:   &res[0].CurrencyID,
		CurrencyName: &res[0].CurrencyName,
		CategoryId:   &res[0].CategoryID,
		CategoryName: &res[0].CategoryName,
		TotalPrice:   &res[0].TotalPrice,
		PurchaseDate: timeToObj(res[0].PurchaseDate),
		Products:     &pros,
	})
}
