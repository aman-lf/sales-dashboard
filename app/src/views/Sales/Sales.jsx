import React from 'react';
import { useQuery } from '@tanstack/react-query';

import Table from '../../components/Table/Table';
import Loading from '../../components/Loading/Loading';

import './sales.scss';

const Sales = ({ title, api, columns }) => {
  const { isPending, error, data } = useQuery({
    queryKey: ['sales'],
    queryFn: () => fetch(`${import.meta.env.VITE_API_URL}${api}`).then((res) => res.json()),
  });

  console.log(import.meta.env.VITE_API_URL, data);

  return (
    <div className='sales'>
      <h2 className='sales__title'>{title}</h2>
      <div>{isPending ? <Loading /> : <Table data={data.data} columns={columns} />}</div>
    </div>
  );
};

export default Sales;
