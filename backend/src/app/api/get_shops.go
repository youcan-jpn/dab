package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) GetShops(ctx echo.Context) error {
	txn, err := a.dbClient.BeginTx(a.ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})
	if err != nil {
		return err
	}

	ss, err := service.GetShops(a.ctx, txn)
	if err != nil {
		return err
	}

	sas := make([]oapi.Shop, 0, len(ss))
	for _, s := range ss {
		sa := oapi.Shop{
			ShopId:   &s.ShopID,
			ShopName: &s.ShopName,
		}
		sas = append(sas, sa)
	}

	return ctx.JSON(http.StatusOK, oapi.Shops{Shops: &sas})
}
