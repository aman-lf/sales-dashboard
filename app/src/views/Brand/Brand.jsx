import Sales from '@/views/Sales';

import { api } from '@/constants/api';
import { title } from '@/constants/title';

const Brand = () => {
  const columns = [
    { key: 'brand_name', label: 'Brand Name' },
    { key: 'most_sold_product', label: 'Most Sold Product' },
    { key: 'total_quantity_sold', label: 'Total Quantity Sold' },
    { key: 'total_revenue', label: 'Total Revenue' },
    { key: 'total_profit', label: 'Total Profit' },
  ];

  return <Sales title={title.SALES_BRAND} api={api.SALES_BRAND} columns={columns} />;
};

export default Brand;
