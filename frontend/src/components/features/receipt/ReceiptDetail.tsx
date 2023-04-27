import type { Receipt } from "#/gen/models";
import { ProductCardList } from "../product/ProductCardList";

interface Prop {
  receipt: Receipt | undefined;
}

export const ReceiptDetail: React.FC<Prop> = (prop: Prop) => {
  const { receipt } = prop;
  const products = receipt?.products;
  const currencyName = receipt?.currency_name;

  if (typeof products !== "undefined" && typeof currencyName !== "undefined") {
    return (
      <>
        <h2>Summary</h2>
        <p>To Be Added</p>
        <h2>Products</h2>
        <ProductCardList products={products} currencyName={currencyName} />
      </>
    );
  }
  return <></>;
};
