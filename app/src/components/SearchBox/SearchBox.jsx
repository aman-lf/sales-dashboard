import { useState } from 'react';

import SearchIcon from '@images/search_icon.svg';

import './searchbox.scss';

const SearchBox = ({ onClick }) => {
  const [text, setText] = useState('');

  return (
    <div className='search-container'>
      <img src={SearchIcon} alt='search icon' className='search-icon' />
      <input
        type='text'
        placeholder='Search...'
        className='search-input'
        value={text}
        onChange={(e) => {
          setText(e.target.value);
        }}
      />
      <button type='submit' className='search-button' onClick={() => onClick(text)}>
        Search
      </button>
    </div>
  );
};

export default SearchBox;
