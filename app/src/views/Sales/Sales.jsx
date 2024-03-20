import React, { useState } from 'react';
import { useQuery } from '@tanstack/react-query';

import Table from '../../components/Table/Table';
import Loading from '../../components/Loading/Loading';
import SearchBox from '../../components/SearchBox/SearchBox';
import Pagination from '../../components/Pagination/Pagination';

import { defaultPerPage } from '../../constants/table';

import { fetchData } from '../../utils/fetch';

import './sales.scss';

const Sales = ({ title, api, columns }) => {
  const [params, setParams] = useState({
    perPage: defaultPerPage,
    page: 1,
    sortBy: '',
    sortOrder: '1',
    searchText: '',
  });

  const { isPending, error, data } = useQuery({
    queryKey: ['sales', params],
    queryFn: () => fetchData(api, params),
  });

  const onSearch = (searchText) => {
    setParams({ ...params, searchText });
  };

  const onPageSizeChange = (perPage) => {
    setParams({ ...params, perPage });
  };

  const onPageChange = (page) => {
    setParams({ ...params, page });
  };

  const onSort = (sortBy, sortOrder) => {
    setParams({ ...params, sortBy, sortOrder });
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
            <Table
              data={data.data}
              columns={columns}
              sortBy={data.meta.sortBy}
              sortOrder={data.meta.sortOrder}
              onSort={onSort}
            />
            <Pagination
              currentPage={params.page}
              totalPages={data.meta.totalPage}
              onPageChange={onPageChange}
              pageSize={params.perPage}
              onPageSizeChange={onPageSizeChange}
            />
          </>
        )}
      </div>
    </div>
  );
};

export default Sales;
