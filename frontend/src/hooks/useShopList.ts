import { atom, useRecoilState } from "recoil";

import type { Shop, ShopId, ShopName } from "#/gen/models";
import {
  usePostShops,
  useGetShops,
  getGetShopsQueryKey,
} from "#/gen/queries/shops/shops";
import { useQueryClient } from "@tanstack/react-query";

const shopListState = atom<Shop[]>({
  key: "shopListState",
  default: [],
});

export const useShopList = () => {
  const queryClient = useQueryClient();
  const postShop = usePostShops();
  const [shopList, setShopList] = useRecoilState(shopListState);
  const { data: res, refetch } = useGetShops();


  const removeById = (shopId: ShopId) => {
    const newShopList = shopList.filter((s) => s.shop_id !== shopId);
    setShopList(newShopList);
  };

  const add = async (shopName: ShopName) => {
    postShop.mutate({ data: { shop_name: shopName } });
    await queryClient.invalidateQueries({ queryKey: [getGetShopsQueryKey()] });
    await refetch();

    if (typeof res?.data.shops !== "undefined") {
      setShopList(res.data.shops);
    }
  };

  return { shopList, setShopList, removeById, add };
};
