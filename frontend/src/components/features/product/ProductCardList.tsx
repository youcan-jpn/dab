import React from "react";
import Box from "@mui/material/Box";

import { ProductCard } from "./ProductCard";
import type { Product, CurrencyName } from "#/gen/models";

interface IProps {
  products: Product[]
  currencyName: CurrencyName
}

export const ProductCardList: React.FC<IProps> = (props: IProps) => {
  const { products, currencyName } = props;

  return (
    <Box>
      {products.map((product: Product) =>
        <ProductCard
          key={product.product_id}
          productId={product.product_id}
          productName={product.product_name}
          price={product.price}
          currencyName={currencyName}
        />
      )}
    </Box>
  )
}