import { Route, Routes } from 'react-router-dom';
import { QueryClientProvider } from '@tanstack/react-query';

import Brand from './views/Brand';
import Product from './views/Products';
import Dashboard from './views/Dashboard';
import AppLayout from './components/Layout/AppLayout';

import { PATHS } from './constants/routes';
import { queryClient } from './config/react-query.config';

import './assets/sass/main.scss';

function App() {
  return (
    <div className='App'>
      <QueryClientProvider client={queryClient}>
        <Routes>
          <Route path={PATHS.HOME_PATH} element={<AppLayout />}>
            <Route index element={<Dashboard />} />
            <Route path={PATHS.PRODUCTS_PATH} element={<Product />} />
            <Route path={PATHS.BRAND_PATH} element={<Brand />} />
          </Route>
        </Routes>
      </QueryClientProvider>
    </div>
  );
}

export default App;
