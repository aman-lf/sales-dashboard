import React from 'react';
import { useQuery } from '@tanstack/react-query';

import Card from '../../components/Card/Card';
import Loading from '../../components/Loading/Loading';

import { fetchData } from '../../utils/fetch';

import './dashboard.scss';

const Dashboard = () => {
  const api = '/api/dashboard';
  const { isPending, error, data } = useQuery({
    queryKey: ['sales'],
    queryFn: () => fetchData(api),
  });

  return (
    <div className='dashboard'>
      <h2 className='dashboard__title'>Dashboard</h2>
      {isPending ? (
        <Loading />
      ) : (
        <div className='card-container'>
          <Card title='Most Profitable Product' description={data.data.most_profitable_product} />
          <Card title='Least Profitable Product' description={data.data.least_profitable_product} />
          <Card title='Hightest sales date' description={data.data.highest_sales_date} />
          <Card title='Least sales date' description={data.data.least_sales_date} />
        </div>
      )}
    </div>
  );
};

export default Dashboard;
