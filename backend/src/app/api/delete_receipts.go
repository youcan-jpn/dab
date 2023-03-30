package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) DeleteReceiptsReceiptId(ctx echo.Context, receiptId oapi.ReceiptId) error {
	txn, err := a.dbClient.BeginTx(a.ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = service.DeleteReceipt(a.ctx, txn, receiptId)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err = txn.Commit(); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.String(http.StatusOK, "The resource was deleted successefully")
}
