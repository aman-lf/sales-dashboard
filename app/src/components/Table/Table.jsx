import { useState } from 'react';

import './table.scss';

const HeaderCell = ({ column, sorting, sortTable }) => {
  const isDescSorting = sorting.column === column.key && sorting.order === 'desc';
  const isAscSorting = sorting.column === column.key && sorting.order === 'asc';
  const futureSortingOrder = isDescSorting ? 'asc' : 'desc';
  return (
    <th key={column.key} className='cell' onClick={() => sortTable({ column: column.key, order: futureSortingOrder })}>
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
      {entries.map((entry) => (
        <tr className='row' key={entry.id}>
          {columns.map((column) => (
            <td key={column.key} className='cell'>
              {entry[column.key]}
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
      <table className='table'>
        <TableHeader columns={columns} sorting={sorting} sortTable={sortTable} />
        <TableBody entries={data} columns={columns} />
      </table>
    </div>
  );
};

export default Table;
