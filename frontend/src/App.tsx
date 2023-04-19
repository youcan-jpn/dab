import axios from 'axios';
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query';

import { SamplePage } from "#/components/pages/Sample.page";
import "./App.css";

axios.defaults.baseURL = "http://localhost:1323/api";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <h1>Hello</h1>
      <h2>Here is DAB page</h2>
      <SamplePage />
    </QueryClientProvider>
  );
}

export default App;
