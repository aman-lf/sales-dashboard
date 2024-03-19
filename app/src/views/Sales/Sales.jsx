import React from 'react';
import { useQuery } from '@tanstack/react-query';

import Table from '../../components/Table/Table';
import Loading from '../../components/Loading/Loading';
import SearchBox from '../../components/SearchBox/SearchBox';

import './sales.scss';
import Pagination from '../../components/Pagination/Pagination';

const Sales = ({ title, api, columns }) => {
  const { isPending, error, data } = useQuery({
    queryKey: ['sales'],
    queryFn: () => fetch(`${import.meta.env.VITE_API_URL}${api}`).then((res) => res.json()),
  });

  const onSearch = (text) => {
    console.log(text);
  };

  if (error) {
    console.log(error.message);
  }

  return (
    <div className='sales'>
      <h2 className='sales__title'>{title}</h2>
      <div>
        <SearchBox onClick={onSearch} />
      </div>
      <div>
        {isPending ? (
          <Loading />
        ) : (
          <>
            <Table data={data.data} columns={columns} />
            <Pagination currentPage={1} totalPages={5} onPageChange='' pageSize={20} onPageSizeChange='' />
          </>
        )}
      </div>
    </div>
  );
};

export default Sales;
