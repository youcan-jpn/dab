package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) PatchCurrenciesCurrencyId(ctx echo.Context, currencyId oapi.CurrencyId) error {
	var cu oapi.PatchCurrenciesCurrencyIdJSONRequestBody
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
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = service.PatchCurrency(a.ctx, txn, currencyId, *cu.CurrencyName, *cu.InYen)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res, err := service.GetOneCurrencyByID(a.ctx, txn, currencyId)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err = txn.Commit(); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	ma := timeToString(res.ModifiedAt)
	ca := timeToString(res.CreatedAt)
	return ctx.JSON(http.StatusOK, oapi.Currency{
		CurrencyId:   &res.CurrencyID,
		CurrencyName: &res.CurrencyName,
		InYen:        &res.InYen,
		ModifiedAt:   &ma,
		CreatedAt:    &ca,
	})
}
