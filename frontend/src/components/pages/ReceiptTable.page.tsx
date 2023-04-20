import {useEffect} from 'react';
import { Box } from "@mui/material";

import type { ReceiptConditionsBody } from "#/gen/models";
import { usePostReceiptsSearch } from "#/gen/queries/receipts/receipts";
import { ReceiptTable } from '../features/receipt/ReceiptTable';

export const ReceiptTablePage = () => {
  const cs: ReceiptConditionsBody = {
    shop_id: 0,
    currency_id: 0,
    category_id: 0,
    min_price: 0,
    max_price: 0,
    purchased_since: {
      year: 0,
      month: 0,
      day: 0,
    },
    purchased_until: {
      year: 0,
      month: 0,
      day: 0,
    },
  };
  const {data: res, refetch} = usePostReceiptsSearch(cs);

  useEffect(() => {
    setTimeout(() => {
      void refetch();
    }, 10000);
  }, [refetch]);

  return (
    <Box component="main">
      <ReceiptTable receipts={res?.data.receipts} />
    </Box>
  );
};
