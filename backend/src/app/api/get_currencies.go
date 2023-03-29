package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) GetCurrencies(ctx echo.Context) error {
	txn, err := a.dbClient.BeginTx(a.ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})
	if err != nil {
		return err
	}

	cus, err := service.GetCurrencies(a.ctx, txn)
	if err != nil {
		return err
	}

	cs := make([]oapi.Currency, 0, len(cus))
	for _, c := range cus {
		iy := float32(c.InYen)
		ma := timeToString(c.ModifiedAt)
		cat := timeToString(c.CreatedAt)
		ca := oapi.Currency{
			CurrencyId: &c.CurrencyID,
			CurrencyName: &c.CurrencyName,
			InYen: &iy,
			ModifiedAt: &ma,
			CreatedAt: &cat,
		}
		cs = append(cs, ca)
	}

	return ctx.JSON(http.StatusOK, oapi.Currencies{Currencies: &cs})
}
