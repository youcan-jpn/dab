import { Routes, Route } from 'react-router-dom';
import { Box } from '@mui/material';

import { HomePage } from './components/pages/Home.page';
import { SamplePage } from "#/components/pages/Sample.page";
import { ReceiptTablePage } from '#/components/pages/ReceiptTable.page';
import { ReceiptDetailPage } from '#/components/pages/ReceiptDetail.page';
import { MenuBar } from '#/components/ui/MenuBar';

import "./App.css";

function App() {
  return (
    <Box className="App">
      <MenuBar />
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/sample" element={<SamplePage />} />
        <Route path="/receipts" element={<ReceiptTablePage />} />
        <Route path="receipts/:receipt_id" element={<ReceiptDetailPage />} />
      </Routes>
    </Box>
  );
}

export default App;
