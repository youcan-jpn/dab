package api

import (
	"context"
	"encoding/json"

	// "encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/testing/database"
)

var shopID1 = 1
var shopID2 = 2
var shopName1 = "sample_shop_1"
var shopName2 = "sample_shop_2"
var modifiedAt1 = "2011-01-01 01:01:01"
var createdAt1 = "2011-01-01 01:00:00"
var modifiedAt2 = "2011-01-02 02:02:02"
var createdAt2 = "2011-01-02 02:00:00"

var wantRes1 = []oapi.Shop{
	{
		ShopId: &shopID1,
		ShopName: &shopName1,
		ModifiedAt: &modifiedAt1,
		CreatedAt: &createdAt1,
	},
	{
		ShopId: &shopID2,
		ShopName: &shopName2,
		ModifiedAt: &modifiedAt2,
		CreatedAt: &createdAt2,
	},
}

var getShopsCases = []struct {
	name string
	reqPath string
	reqBody string
	wantSC int  // status code
	wantRes oapi.Shops
	wantErr bool
} {
	{
		name: "successful case",
		reqPath: "/shops",
		reqBody: "",
		wantSC: http.StatusOK,
		wantRes: oapi.Shops{
			Shops: &wantRes1,
		},
	},
}

func TestAPI_GetShops(t *testing.T) {
	t.Parallel()

	e := echo.New()

	ctx := context.Background()

	for _, tc := range getShopsCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.reqPath, strings.NewReader(tc.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ectx := e.NewContext(req, rec)
			db, err := database.NewFakeDB()
			if err != nil {
				t.Error("initializing db failed")
			}

			a := NewAPI(ctx, db, "", nil)
			err = a.GetShops(ectx)
			if err != nil {
				t.Error("Binding failed")
			}
			var res oapi.Shops
			err = json.Unmarshal(rec.Body.Bytes(), &res)
			if err != nil {
				t.Errorf("Marshaling failed")
			}

			assert.Equal(t, tc.wantSC, rec.Code)
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}