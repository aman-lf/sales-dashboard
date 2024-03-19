import React from 'react';

import Sales from '../Sales';

const Brand = () => {
  const api = '/api/sale-brand';
  const columns = [
    { key: 'brand_name', label: 'Brand Name' },
    { key: 'most_sold_product', label: 'Most Sold Product' },
    { key: 'total_quantity_sold', label: 'Total Quantity Sold' },
    { key: 'total_revenue', label: 'Total Revenue' },
    { key: 'total_profit', label: 'Total Profit' },
  ];

  const searchColumns = ['brand_name', 'most_sold_product'];

  return <Sales title={'Sales By Product'} api={api} columns={columns} searchColumns={searchColumns} />;
};

export default Brand;
