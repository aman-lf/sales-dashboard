import { useState } from 'react';

import './table.scss';

const HeaderCell = ({ column, sorting, sortTable }) => {
  const isDescSorting = sorting.column === column && sorting.order === 'desc';
  const isAscSorting = sorting.column === column && sorting.order === 'asc';
  const futureSortingOrder = isDescSorting ? 'asc' : 'desc';
  return (
    <th key={column} className='users-table-cell' onClick={() => sortTable({ column, order: futureSortingOrder })}>
      {column}
      {isDescSorting && <span>▼</span>}
      {isAscSorting && <span>▲</span>}
    </th>
  );
};

const TableHeader = ({ columns, sorting, sortTable }) => {
  return (
    <thead>
      <tr>
        {columns.map((column) => (
          <HeaderCell column={column} sorting={sorting} key={column} sortTable={sortTable} />
        ))}
      </tr>
    </thead>
  );
};

const TableBody = ({ entries, columns }) => {
  return (
    <tbody>
      {entries.map((entry) => (
        <tr key={entry.id}>
          {columns.map((column) => (
            <td key={column} className='users-table-cell'>
              {entry[column]}
            </td>
          ))}
        </tr>
      ))}
    </tbody>
  );
};

const Table = ({ data, columns }) => {
  const [sorting, setSorting] = useState({ column: 'id', order: 'asc' });
  const sortTable = (newSorting) => {
    setSorting(newSorting);
  };

  return (
    <div>
      <table className='users-table'>
        <TableHeader columns={columns} sorting={sorting} sortTable={sortTable} />
        <TableBody entries={data} columns={columns} />
      </table>
    </div>
  );
};

export default Table;
