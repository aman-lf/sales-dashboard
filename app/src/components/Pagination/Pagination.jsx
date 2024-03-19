import React from 'react';

import './pagination.scss';

const Pagination = ({ currentPage, totalPages, onPageChange, pageSize, onPageSizeChange }) => {
  const handlePreviousPage = () => {
    if (currentPage > 1) {
      onPageChange(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      onPageChange(currentPage + 1);
    }
  };

  const handlePageSizeChange = (e) => {
    const newSize = parseInt(e.target.value, 10);
    onPageSizeChange(newSize);
  };

  return (
    <div className='pagination-container'>
      <button onClick={handlePreviousPage} disabled={currentPage === 1} className='pagination-button left'>
        &#9664;
      </button>
      <span className='pagination-text'>
        Page {currentPage} of {totalPages}
      </span>
      <button onClick={handleNextPage} disabled={currentPage === totalPages} className='pagination-button right'>
        &#9654;
      </button>
      <select value={pageSize} onChange={handlePageSizeChange} className='page-select'>
        <option value='5'>5 per page</option>
        <option value='10'>10 per page</option>
        <option value='15'>15 per page</option>
        <option value='20'>20 per page</option>
      </select>
    </div>
  );
};

export default Pagination;
