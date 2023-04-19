import { Typography, Box } from "@mui/material";

import { useGetShops } from "#/gen/queries/shops/shops";
import type { Shop } from "#/gen/models";

export const SamplePage = () => {
  const {data: res} = useGetShops();
  console.log(res?.data);
  return (
    <Box component="main">
      <Typography variant="h3">Sample</Typography>
      <Typography variant="body1">Sample page of this web app.</Typography>
      {res?.data.shops?.map((shop: Shop) => (
        <p key={shop.shop_id}>{shop.shop_id}: {shop.shop_name}</p>
      ))}
    </Box>
  );
};
