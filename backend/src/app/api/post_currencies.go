package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) PostCurrencies(ctx echo.Context) error {
	var cu oapi.NewOrUpdatedCurrency
	err := ctx.Bind(&cu)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	txn, err := a.dbClient.BeginTx(a.ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	defer txn.Rollback()
	if err != nil {
		return err
	}

	r, err := service.PostCurrency(a.ctx, txn, *cu.CurrencyName, float32(*cu.InYen))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	lid, err := r.LastInsertId()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res, err := service.GetOneCurrencyByID(a.ctx, txn, int(lid))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err = txn.Commit(); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	iy := float32(res.InYen)
	ma := timeToString(res.ModifiedAt)
	ca := timeToString(res.CreatedAt)

	return ctx.JSON(http.StatusOK, oapi.Currency{
		CurrencyId: &res.CurrencyID,
		CurrencyName: &res.CurrencyName,
		InYen: &iy,
		ModifiedAt: &ma,
		CreatedAt: &ca,
	})
}
