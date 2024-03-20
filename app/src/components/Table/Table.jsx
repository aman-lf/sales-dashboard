import { useState } from 'react';

import './table.scss';

const HeaderCell = ({ column, sorting, sortTable }) => {
  const isDescSorting = sorting.column === column.key && sorting.order === -1;
  const isAscSorting = sorting.column === column.key && sorting.order === 1;
  const futureSortingOrder = isDescSorting ? 1 : -1;
  return (
    <th key={column.key} className='cell' onClick={() => sortTable(column.key, futureSortingOrder)}>
      {column.label}
      {isDescSorting && <span>▼</span>}
      {isAscSorting && <span>▲</span>}
    </th>
  );
};

const TableHeader = ({ columns, sorting, sortTable }) => {
  return (
    <thead>
      <tr className='row header'>
        {columns.map((column) => (
          <HeaderCell column={column} sorting={sorting} key={column.key} sortTable={sortTable} />
        ))}
      </tr>
    </thead>
  );
};

const TableBody = ({ entries, columns }) => {
  return (
    <tbody>
      {entries.length == 0 ? (
        <tr>
          <td className='no-data' colSpan={columns.length}>
            No Data Available
          </td>
        </tr>
      ) : (
        <>
          {entries.map((entry, index) => (
            <tr className='row' key={`row${index}`}>
              {columns.map((column) => (
                <td key={column.key} className='cell'>
                  {entry[column.key]}
                </td>
              ))}
            </tr>
          ))}
        </>
      )}
    </tbody>
  );
};

const Table = ({ data, columns, sortBy, sortOrder, onSort }) => {
  const sorting = { column: sortBy, order: sortOrder };

  return (
    <div>
      <table className='table'>
        <TableHeader columns={columns} sorting={sorting} sortTable={onSort} />
        <TableBody entries={data} columns={columns} />
      </table>
    </div>
  );
};

export default Table;
