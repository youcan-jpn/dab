import React from "react";
import { useNavigate } from "react-router-dom";
import Divider from "@mui/material/Divider";
import Paper from "@mui/material/Paper";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import ListItemText from "@mui/material/ListItemText";
import ListItemIcon from "@mui/material/ListItemIcon";
import LoupeIcon from "@mui/icons-material/Loupe";
import DeleteIcon from "@mui/icons-material/Delete";

import {
  useDeleteReceiptsReceiptId,
  usePostReceiptsSearch,
} from "#/gen/queries/receipts/receipts";
// import DeleteErrorDialog from './DeleteErrorDialog';
import type { ReceiptId, ReceiptConditionsBody } from "#/gen/models";
import axios from "axios";

interface IProps {
  receiptId: ReceiptId;
  anchorEl: HTMLElement | null;
  open: boolean;
  onClose: () => void;
}

const ReceiptMenu: React.FC<IProps> = (props) => {
  const { receiptId, anchorEl, open, onClose } = props;
  const deleteFn = useDeleteReceiptsReceiptId();
  const navigate = useNavigate();
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
  const { refetch } = usePostReceiptsSearch(cs);

  const HandleClick = () => {
    navigate(`/receipts/${receiptId}`);
  };

  const HandleDelete = () => {
    try {
      void deleteFn.mutateAsync({ receiptId });
      console.log("DELETED!");
      void refetch();
    } catch (e) {
      if (axios.isAxiosError(e)) {
        const statusCode = e.response?.status;
        const data = e.response?.data;
        const errMsg = data?.message;
        console.log(statusCode);
        console.log(errMsg);
      }
    }
  };

  return (
    <Paper sx={{ width: 320, maxWidth: "100%" }}>
      <Menu open={open} anchorEl={anchorEl} onClose={onClose}>
        <MenuItem onClick={HandleClick}>
          <ListItemIcon>
            <LoupeIcon fontSize="small" />
          </ListItemIcon>
          <ListItemText>View Detail</ListItemText>
        </MenuItem>
        <Divider />
        <MenuItem onClick={HandleDelete}>
          <ListItemIcon>
            <DeleteIcon fontSize="small" />
          </ListItemIcon>
          <ListItemText>Delete</ListItemText>
        </MenuItem>
      </Menu>
    </Paper>
  );
};

export default ReceiptMenu;
