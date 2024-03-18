import React from 'react';

import Sales from '../Sales/Sales';

const Product = () => {
  const api = '/api/sale-product';
  const columns = [
    { key: 'id', label: 'ID' },
    { key: 'product_name', label: 'Product Name' },
    { key: 'brand_name', label: 'Brand Name' },
    { key: 'category', label: 'Category' },
    { key: 'total_quantity_sold', label: 'Total Quantity Sold' },
    { key: 'total_revenue', label: 'Total Revenue' },
    { key: 'total_profit', label: 'Total Profit' },
  ];

  return <Sales title={'Sales By Product'} api={api} columns={columns} />;
};

export default Product;
