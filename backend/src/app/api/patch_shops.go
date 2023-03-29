package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) PatchShopsShopId(ctx echo.Context, shopId oapi.ShopId) error {
	var shop oapi.Shop
	err := ctx.Bind(&shop)
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

	err = service.PatchShop(a.ctx, txn, shopId, *shop.ShopName)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res, err := service.GetOneShopByID(a.ctx, txn, shopId)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	ma := timeToString(res.ModifiedAt)
	ca := timeToString(res.CreatedAt)

	if err = txn.Commit(); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, oapi.Shop{
		ShopId:     &res.ShopID,
		ShopName:   &res.ShopName,
		ModifiedAt: &ma,
		CreatedAt:  &ca,
	})
}
