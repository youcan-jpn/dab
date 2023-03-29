package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/youcan-jpn/dab/backend/dao"
	"github.com/youcan-jpn/dab/backend/gen/daocore"
)

type Receipt struct {
	ReceiptID     int
	ShopID        int
	CurrencyID    int
	CategoryID    int
	TotalPrice    float32
	purchase_date *time.Time
	Products      []*Product
}

type Product struct {
	ProductID   int
	ProductName string
	Price       float32
}

func (p Product) ToDAOProduct(rid int) (*daocore.Product, error) {
	if p.ProductID <= 0 {
		return &daocore.Product{}, fmt.Errorf("'product_id' has to be positive")
	}
	if p.ProductName == "" {
		return &daocore.Product{}, fmt.Errorf("'product_name' has to be set")
	}
	if p.Price <= 0 {
		return &daocore.Product{}, fmt.Errorf("'price' has to be positive")
	}
	return &daocore.Product{
		ReceiptID:   rid,
		ProductID:   p.ProductID,
		ProductName: p.ProductName,
		Price:       p.Price,
	}, nil
}

// Receipts
func PostReceipt(ctx context.Context, txn *sql.Tx, shop_id int, currency_id int, category_id int, purchase_date *time.Time, total_price float32, products []*Product) error {
	if shop_id <= 0 {
		return fmt.Errorf("'shop_id' has to be positive")
	}
	if currency_id <= 0 {
		return fmt.Errorf("'currency_id' has to be positive")
	}
	if category_id <= 0 {
		return fmt.Errorf("'category_id' has to be positive")
	}
	if purchase_date == nil {
		return fmt.Errorf("'purchase_date' has to be not nil")
	}
	if total_price <= 0 {
		return fmt.Errorf("'total_price' has to be positive")
	}
	if products == nil {
		return fmt.Errorf("'products' has to be not nil")
	}

	r := daocore.Receipt{
		ReceiptID:    0,
		ShopID:       shop_id,
		CurrencyID:   currency_id,
		CategoryID:   category_id,
		TotalPrice:   total_price,
		PurchaseDate: purchase_date,
	}

	res, err := dao.InsertOneReceiptReturningResult(ctx, txn, &r)
	if err != nil {
		return err
	}

	rid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	ps := make([]*daocore.Product, 0, len(products))
	for _, p := range products {
		pp, err := p.ToDAOProduct(int(rid))
		if err != nil {
			return err
		}
		ps = append(ps, pp)
	}

	err = daocore.InsertProduct(ctx, txn, ps)
	return err
}

func PostReceiptsSearch(ctx context.Context, txn *sql.Tx, receipt_id int, shop_id int, currency_id int, category_id int, min_price float32, max_price float32, since *time.Time, until *time.Time) ([]*dao.ReceiptDetail, error) {
	return dao.SelectReceiptDetailsByConditions(ctx, txn, receipt_id, shop_id, currency_id, category_id, min_price, max_price, since, until)
}

func GetReceipts(ctx context.Context, txn *sql.Tx, receipt_id int) ([]*dao.ReceiptDetail, error) {
	if receipt_id > 0 {
		return dao.SelectReceiptDetailsByConditions(ctx, txn, receipt_id, 0, 0, 0, 0, 0, nil, nil)
	}
	return nil, fmt.Errorf("receipt_id has to be integer (> 0)")
}

func PatchReceipt(ctx context.Context, txn *sql.Tx, rec Receipt) error {
	r := daocore.Receipt{
		ReceiptID:    rec.ReceiptID,
		ShopID:       rec.ShopID,
		CurrencyID:   rec.CurrencyID,
		TotalPrice:   rec.TotalPrice,
		PurchaseDate: rec.purchase_date,
	}
	err := daocore.UpsertReceipt(ctx, txn, r)
	if err != nil {
		return err
	}

	err = daocore.DeleteProductByReceiptID(ctx, txn, r.ReceiptID)
	if err != nil {
		return err
	}

	ps := make([]*daocore.Product, 0, len(rec.Products))
	for _, p := range rec.Products {
		pp, err := p.ToDAOProduct(r.ReceiptID)
		if err != nil {
			return err
		}
		ps = append(ps, pp)
	}

	return daocore.InsertProduct(ctx, txn, ps)
}

func DeleteReceipt(ctx context.Context, txn *sql.Tx, receipt_id int) error {
	err := daocore.DeleteProductByReceiptID(ctx, txn, receipt_id)
	if err != nil {
		return err
	}
	return daocore.DeleteOneReceiptByReceiptID(ctx, txn, receipt_id)
}

// Shops
func GetOneShopByID(ctx context.Context, txn *sql.Tx, shop_id int) (daocore.Shop, error) {
	return daocore.SelectOneShopByShopID(ctx, txn, &shop_id)
}
func GetShops(ctx context.Context, txn *sql.Tx) ([]*daocore.Shop, error) {
	return daocore.SelectShopByShopName(ctx, txn, nil)
}

func PostShop(ctx context.Context, txn *sql.Tx, shop_name string) (sql.Result, error) {
	return dao.InsertOneShopReturningResult(ctx, txn, &daocore.Shop{ShopID: 0, ShopName: shop_name})
}

func PatchShop(ctx context.Context, txn *sql.Tx, shop_id int, shop_name string) error {
	if shop_id > 0 && len(shop_name) > 0 {
		return daocore.UpsertShop(ctx, txn, daocore.Shop{ShopID: shop_id, ShopName: shop_name})
	}
	return fmt.Errorf("'shop_id' and 'shop_name' have to be set")
}

// Currencies
func GetCurrencies(ctx context.Context, txn *sql.Tx) ([]*daocore.Currency, error) {
	return daocore.SelectCurrencyByInYen(ctx, txn, nil)
}

func PostCurrency(ctx context.Context, txn *sql.Tx, currency_name string, in_yen float32) error {
	return daocore.InsertCurrency(ctx, txn, []*daocore.Currency{{CurrencyID: 0, CurrencyName: currency_name, InYen: in_yen}})
}

func PatchCurrency(ctx context.Context, txn *sql.Tx, currency_id int, currency_name string, in_yen float32) error {
	if currency_id > 0 && len(currency_name) == 3 {
		return daocore.UpsertCurrency(ctx, txn, daocore.Currency{CurrencyID: currency_id, CurrencyName: currency_name, InYen: in_yen})
	}
	return fmt.Errorf("'currency_id' and 'currency_name' (len=3) have to be set")
}

// Categories
func GetCategories(ctx context.Context, txn *sql.Tx) ([]*daocore.Category, error) {
	return daocore.SelectCategoryByCategoryName(ctx, txn, nil)
}

func PostCategory(ctx context.Context, txn *sql.Tx, category_name string) error {
	return daocore.InsertCategory(ctx, txn, []*daocore.Category{{CategoryID: 0, CategoryName: category_name}})
}

func PatchCategory(ctx context.Context, txn *sql.Tx, category_id int, category_name string) error {
	if category_id > 0 && len(category_name) > 0{
		return daocore.UpsertCategory(ctx, txn, daocore.Category{CategoryID: category_id, CategoryName: category_name})
	}
	return fmt.Errorf("'category_id' and 'category_name' have to be set")
}
