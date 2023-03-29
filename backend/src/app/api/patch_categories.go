package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) PatchCategoriesCategoryId(ctx echo.Context, categoryId oapi.CategoryId) error {
	var cate oapi.Category
	err := ctx.Bind(&cate)
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

	err = service.PatchCategory(a.ctx, txn, categoryId, *cate.CategoryName)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res, err := service.GetOneCategoryByID(a.ctx, txn, categoryId)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err = txn.Commit(); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, oapi.Category{
		CategoryId:     &res.CategoryID,
		CategoryName:   &res.CategoryName,
	})
}
