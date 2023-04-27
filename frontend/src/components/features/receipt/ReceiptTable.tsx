import { useState } from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import IconButton from "@mui/material/IconButton";

import ReceiptMenu from "./ReceiptMenu";
import type { Receipt, ReceiptId } from "#/gen/models";

import { displayPrice } from "#/components/functional/DisplayPrice";
import { displayDate } from "#/components/functional/DisplayDate";

interface Prop {
  receipts: Receipt[] | undefined;
}

export const ReceiptTable: React.FC<Prop> = (prop: Prop) => {
  const { receipts } = prop;
  const [openMenu, setOpenMenu] = useState<boolean>(false);
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);
  const [tgtReceiptId, setTgtReceiptId] = useState<ReceiptId>(0);

  const HandleOpen = (e: any, receiptId: ReceiptId | undefined) => {
    if (typeof receiptId === "undefined") {
      setTgtReceiptId(0);
    } else {
      setTgtReceiptId(receiptId);
    }
    setAnchorEl(e.currentTarget);
    setOpenMenu(true);
  };

  const HandleClose = () => {
    setOpenMenu(false);
    setAnchorEl(null);
    setTgtReceiptId(0);
  };

  return (
    <>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">Receipt ID</TableCell>
              <TableCell align="center">Shop Name</TableCell>
              <TableCell align="center">Total Price</TableCell>
              <TableCell align="center">Purchase Date</TableCell>
              <TableCell align="center">Options</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {receipts?.map((row: Receipt) => (
              <TableRow key={row.receipt_id}>
                <TableCell component="th" scope="row" align="center">
                  {row.receipt_id}
                </TableCell>
                <TableCell align="center">{row.shop_name}</TableCell>
                <TableCell align="center">
                  {displayPrice(row.total_price, row.currency_name)}
                </TableCell>
                <TableCell align="center">
                  {displayDate(row.purchase_date)}
                </TableCell>
                <TableCell align="center">
                  <IconButton
                    onClick={(e) => {
                      HandleOpen(e, row.receipt_id);
                    }}
                  >
                    <MoreVertIcon />
                  </IconButton>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <ReceiptMenu
        receiptId={tgtReceiptId}
        anchorEl={anchorEl}
        onClose={HandleClose}
        open={openMenu}
      />
    </>
  );
};
