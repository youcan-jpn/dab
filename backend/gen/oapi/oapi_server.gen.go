// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package oapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// カテゴリ情報を取得
	// (GET /categories)
	GetCategories(ctx echo.Context) error
	// カテゴリ情報を登録
	// (POST /categories)
	PostCategories(ctx echo.Context) error
	// カテゴリ情報を更新
	// (PATCH /categories/{category_id})
	PatchCategoriesCategoryId(ctx echo.Context, categoryId CategoryId) error
	// 通貨情報を取得
	// (GET /currencies)
	GetCurrencies(ctx echo.Context) error
	// 通貨情報を登録
	// (POST /currencies)
	PostCurrencies(ctx echo.Context) error
	// 通貨情報を更新
	// (PATCH /currencies/{currency_id})
	PatchCurrenciesCurrencyId(ctx echo.Context, currencyId CurrencyId) error
	// 新しいレシートを登録
	// (POST /receipts)
	PostReceipts(ctx echo.Context) error
	// 条件を指定して該当するレシートを検索
	// (POST /receipts/search)
	PostReceiptsSearch(ctx echo.Context) error
	// レシートIDを指定して削除
	// (DELETE /receipts/{receipt_id})
	DeleteReceiptsReceiptId(ctx echo.Context, receiptId ReceiptId) error
	// レシートIDを指定して詳細を取得
	// (GET /receipts/{receipt_id})
	GetReceiptsReceiptId(ctx echo.Context, receiptId ReceiptId) error
	// レシートIDを指定して更新
	// (PATCH /receipts/{receipt_id})
	PatchReceiptsReceiptId(ctx echo.Context, receiptId ReceiptId) error
	// 店一覧を取得
	// (GET /shops)
	GetShops(ctx echo.Context) error
	// 新しい店名を登録
	// (POST /shops)
	PostShops(ctx echo.Context) error
	// 店名を更新
	// (PATCH /shops/{shop_id})
	PatchShopsShopId(ctx echo.Context, shopId ShopId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCategories converts echo context to params.
func (w *ServerInterfaceWrapper) GetCategories(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCategories(ctx)
	return err
}

// PostCategories converts echo context to params.
func (w *ServerInterfaceWrapper) PostCategories(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostCategories(ctx)
	return err
}

// PatchCategoriesCategoryId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchCategoriesCategoryId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "category_id" -------------
	var categoryId CategoryId

	err = runtime.BindStyledParameterWithLocation("simple", false, "category_id", runtime.ParamLocationPath, ctx.Param("category_id"), &categoryId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter category_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchCategoriesCategoryId(ctx, categoryId)
	return err
}

// GetCurrencies converts echo context to params.
func (w *ServerInterfaceWrapper) GetCurrencies(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCurrencies(ctx)
	return err
}

// PostCurrencies converts echo context to params.
func (w *ServerInterfaceWrapper) PostCurrencies(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostCurrencies(ctx)
	return err
}

// PatchCurrenciesCurrencyId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchCurrenciesCurrencyId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "currency_id" -------------
	var currencyId CurrencyId

	err = runtime.BindStyledParameterWithLocation("simple", false, "currency_id", runtime.ParamLocationPath, ctx.Param("currency_id"), &currencyId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter currency_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchCurrenciesCurrencyId(ctx, currencyId)
	return err
}

// PostReceipts converts echo context to params.
func (w *ServerInterfaceWrapper) PostReceipts(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostReceipts(ctx)
	return err
}

// PostReceiptsSearch converts echo context to params.
func (w *ServerInterfaceWrapper) PostReceiptsSearch(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostReceiptsSearch(ctx)
	return err
}

// DeleteReceiptsReceiptId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteReceiptsReceiptId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "receipt_id" -------------
	var receiptId ReceiptId

	err = runtime.BindStyledParameterWithLocation("simple", false, "receipt_id", runtime.ParamLocationPath, ctx.Param("receipt_id"), &receiptId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter receipt_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteReceiptsReceiptId(ctx, receiptId)
	return err
}

// GetReceiptsReceiptId converts echo context to params.
func (w *ServerInterfaceWrapper) GetReceiptsReceiptId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "receipt_id" -------------
	var receiptId ReceiptId

	err = runtime.BindStyledParameterWithLocation("simple", false, "receipt_id", runtime.ParamLocationPath, ctx.Param("receipt_id"), &receiptId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter receipt_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReceiptsReceiptId(ctx, receiptId)
	return err
}

// PatchReceiptsReceiptId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchReceiptsReceiptId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "receipt_id" -------------
	var receiptId ReceiptId

	err = runtime.BindStyledParameterWithLocation("simple", false, "receipt_id", runtime.ParamLocationPath, ctx.Param("receipt_id"), &receiptId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter receipt_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchReceiptsReceiptId(ctx, receiptId)
	return err
}

// GetShops converts echo context to params.
func (w *ServerInterfaceWrapper) GetShops(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetShops(ctx)
	return err
}

// PostShops converts echo context to params.
func (w *ServerInterfaceWrapper) PostShops(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostShops(ctx)
	return err
}

// PatchShopsShopId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchShopsShopId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "shop_id" -------------
	var shopId ShopId

	err = runtime.BindStyledParameterWithLocation("simple", false, "shop_id", runtime.ParamLocationPath, ctx.Param("shop_id"), &shopId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shop_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchShopsShopId(ctx, shopId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/categories", wrapper.GetCategories)
	router.POST(baseURL+"/categories", wrapper.PostCategories)
	router.PATCH(baseURL+"/categories/:category_id", wrapper.PatchCategoriesCategoryId)
	router.GET(baseURL+"/currencies", wrapper.GetCurrencies)
	router.POST(baseURL+"/currencies", wrapper.PostCurrencies)
	router.PATCH(baseURL+"/currencies/:currency_id", wrapper.PatchCurrenciesCurrencyId)
	router.POST(baseURL+"/receipts", wrapper.PostReceipts)
	router.POST(baseURL+"/receipts/search", wrapper.PostReceiptsSearch)
	router.DELETE(baseURL+"/receipts/:receipt_id", wrapper.DeleteReceiptsReceiptId)
	router.GET(baseURL+"/receipts/:receipt_id", wrapper.GetReceiptsReceiptId)
	router.PATCH(baseURL+"/receipts/:receipt_id", wrapper.PatchReceiptsReceiptId)
	router.GET(baseURL+"/shops", wrapper.GetShops)
	router.POST(baseURL+"/shops", wrapper.PostShops)
	router.PATCH(baseURL+"/shops/:shop_id", wrapper.PatchShopsShopId)

}
