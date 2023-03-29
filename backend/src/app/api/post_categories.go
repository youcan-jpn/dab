package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/service"
)

func (a *API) PostCategories(ctx echo.Context) error {
	var cate oapi.NewOrUpdatedCategory
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
		return err
	}

	r, err := service.PostCategory(a.ctx, txn, *cate.CategoryName)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	lid, err := r.LastInsertId()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res, err := service.GetOneCategoryByID(a.ctx, txn, int(lid))
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
