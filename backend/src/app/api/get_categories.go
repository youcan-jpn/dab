package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) GetCategories(ctx echo.Context) error {
	txn, err := a.dbClient.BeginTx(a.ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})
	if err != nil {
		return err
	}

	cas, err := service.GetCategories(a.ctx, txn)
	if err != nil {
		return err
	}

	cs := make([]oapi.Category, 0, len(cas))
	for _, c := range cas {
		ca := oapi.Category{
			CategoryId:   &c.CategoryID,
			CategoryName: &c.CategoryName,
		}
		cs = append(cs, ca)
	}

	return ctx.JSON(http.StatusOK, oapi.Categories{Categories: &cs})
}
