/**
 * Generated by orval v6.12.1 🍺
 * Do not edit manually.
 * dab-api
 * 家計簿アプリ用API
 * OpenAPI spec version: 0.0.1
 */
import type { ShopId } from "./shopId";
import type { ShopName } from "./shopName";
import type { Timestamp } from "./timestamp";

/**
 * 店情報
 */
export interface Shop {
  shop_id?: ShopId;
  shop_name?: ShopName;
  modified_at?: Timestamp;
  created_at?: Timestamp;
}
