import { useParams } from "react-router-dom";

import { ReceiptDetail } from "../features/receipt/ReceiptDetail";
import { useGetReceiptsReceiptId } from "#/gen/queries/receipts/receipts";

export const ReceiptDetailPage = () => {
  const params = useParams();
  const receiptId = Number(params.receipt_id);

  const { data: res } = useGetReceiptsReceiptId(receiptId);

  const receipt = res?.data;

  return (
    <>
      <ReceiptDetail receipt={receipt} />
    </>
  );
};
