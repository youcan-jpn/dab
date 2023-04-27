import React, { useState } from "react";
import TextField from "@mui/material/TextField/TextField";
import Button from "@mui/material/Button/Button";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { useShopList } from "#/hooks/useShopList";
import type { ShopName } from "#/gen/models";

export const AddShopInput: React.FC = () => {
  const { add } = useShopList();
  const [shopName, setShopName] = useState<ShopName>("");

  const changeHandler = (e: any) => {
    setShopName(e.target.value);
  };

  const clickHandler = (): void => {
    void add(shopName);
  };

  return (
    <Card sx={{ minWidth: 275 }}>
      <CardContent>
        <Typography sx={{ fontSize: 14 }} color="text.secondary">
          Input new shop&apos;s name:
        </Typography>
        <TextField
          id="add-new-shop"
          label="new shop"
          variant="outlined"
          onChange={changeHandler}
        />
      </CardContent>
      <CardActions>
        <Button
          size="small"
          onClick={() => {
            clickHandler();
          }}
        >
          Register
        </Button>
      </CardActions>
    </Card>
  );
};
