import Sales from '@/views/Sales';

import { api } from '@/constants/api';
import { title } from '@/constants/title';

const Product = () => {
  const columns = [
    { key: 'product_id', label: 'ID' },
    { key: 'product_name', label: 'Product Name' },
    { key: 'brand_name', label: 'Brand Name' },
    { key: 'category', label: 'Category' },
    { key: 'total_quantity_sold', label: 'Total Quantity Sold' },
    { key: 'total_revenue', label: 'Total Revenue' },
    { key: 'total_profit', label: 'Total Profit' },
  ];

  return <Sales title={title.SALES_PRODUCT} api={api.SALES_PRODUCT} columns={columns} />;
};

export default Product;
