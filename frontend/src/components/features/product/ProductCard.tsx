import React from "react";
import { Card, CardContent } from "@mui/material";
import Typography from "@mui/material/Typography";

import { displayPrice } from "#/components/functional/DisplayPrice";
import type { Price } from "#/gen/models";

interface IProps {
  productId: number | undefined;
  productName: string | undefined;
  price: Price | undefined;
  currencyName: string | undefined;
}

export const ProductCard: React.FC<IProps> = (props: IProps) => {
  const { productId, productName, price, currencyName } = props;

  return (
    <Card>
      <CardContent>
        <Typography>product ID: {productId}</Typography>
        <Typography>{productName}</Typography>
        <Typography>{displayPrice(price, currencyName)}</Typography>
      </CardContent>
    </Card>
  );
};
