package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) GetReceiptsReceiptId(ctx echo.Context, receiptId oapi.ReceiptId) error {
	txn, err := a.dbClient.BeginTx(a.ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})
	if err != nil {
		return err
	}

	rs, err := service.GetReceipts(a.ctx, txn, receiptId)
	if err != nil {
		return err
	}
	if len(rs) == 0 {
		return ctx.JSON(http.StatusOK, oapi.Receipt{})
	}

	ps := make([]oapi.Product, 0, len(rs))
	for _, r := range rs {
		p := oapi.Product{
			ProductId:   &r.ProductID,
			ProductName: &r.ProductName,
			Price:       &r.Price,
		}
		ps = append(ps, p)
	}

	return ctx.JSON(http.StatusOK, oapi.Receipt{
		ReceiptId:    &rs[0].ReceiptID,
		ShopId:       &rs[0].ShopID,
		ShopName:     &rs[0].ShopName,
		CurrencyId:   &rs[0].CurrencyID,
		CurrencyName: &rs[0].CurrencyName,
		CategoryId:   &rs[0].CategoryID,
		CategoryName: &rs[0].CategoryName,
		PurchaseDate: timeToObj(rs[0].PurchaseDate),
		TotalPrice:   &rs[0].TotalPrice,
		Products:     &ps,
	})
}
