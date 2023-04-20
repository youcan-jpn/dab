import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

import type { Receipt, Date, Price, CurrencyName } from '#/gen/models';

interface Prop {
  receipts: Receipt[] | undefined
}


export const ReceiptTable: React.FC<Prop> = (prop: Prop) => {
  const { receipts } = prop;

  const displayDate = (date: Date | undefined): string => {
    if (typeof date !== 'undefined') {
      if (typeof date.year === 'number' && typeof date.month === 'number' && typeof date.day === 'number') {
        const dateObject = new Date(`${date.year}-${date.month}-${date.day}`);
        const ret = dateObject.toLocaleDateString();
        return ret;
      }
    }
    return '';
  };

  const displayPrice = (price: Price | undefined, currencyName: CurrencyName | undefined): string => {
    if (typeof price === 'undefined' || typeof currencyName === 'undefined') {
      return '';
    }
    return `${price.toLocaleString()} ${currencyName}`;
  }

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
            </TableRow>
          </TableHead>
          <TableBody>
            {receipts?.map((row: Receipt) => (
              <TableRow key={row.receipt_id}>
                <TableCell component="th" scope="row" align="center">
                  {row.receipt_id}
                </TableCell>
                <TableCell align="center">
                  {row.shop_name}
                </TableCell>
                <TableCell align="center">
                  {displayPrice(row.total_price, row.currency_name)}
                </TableCell>
                <TableCell align="center">
                  {displayDate(row.purchase_date)}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
