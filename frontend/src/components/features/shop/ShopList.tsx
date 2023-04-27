import React, { useEffect } from "react";

import { useGetShops } from "#/gen/queries/shops/shops";
import { useShopList } from "#/hooks/useShopList";
import type { Shop } from "#/gen/models";

export const ShopList: React.FC = () => {
  const { shopList: shops, setShopList } = useShopList();
  const { data: res } = useGetShops();

  useEffect(() => {
    if (typeof res?.data.shops !== "undefined") {
      setShopList(res?.data.shops);
    }
  }, []);

  return (
    <>
      <ul>
        {shops.map((s: Shop) => (
          <li key={s.shop_id}>
            {s.shop_id}: {s.shop_name}
          </li>
        ))}
      </ul>
    </>
  );
};
