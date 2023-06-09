/**
 * Generated by orval v6.12.1 🍺
 * Do not edit manually.
 * dab-api
 * 家計簿アプリ用API
 * OpenAPI spec version: 0.0.1
 */
import type { ReceiptId } from "./receiptId";
import type { ShopId } from "./shopId";
import type { ShopName } from "./shopName";
import type { CurrencyId } from "./currencyId";
import type { CurrencyName } from "./currencyName";
import type { CategoryId } from "./categoryId";
import type { CategoryName } from "./categoryName";
import type { Date } from "./date";
import type { Price } from "./price";
import type { Products } from "./products";

/**
 * レシート一覧の1要素
 */
export interface Receipt {
  receipt_id?: ReceiptId;
  shop_id?: ShopId;
  shop_name?: ShopName;
  currency_id?: CurrencyId;
  currency_name?: CurrencyName;
  category_id?: CategoryId;
  category_name?: CategoryName;
  purchase_date?: Date;
  total_price?: Price;
  products?: Products;
}
