import NotFoundImg from '@images/not_found.webp';

import './notFound.scss';

const NotFound = () => {
  return (
    <div className='not-found'>
      <img src={NotFoundImg} alt='not found' className='not-found-image' />
      <div className='text'>
        <h1>404 - Page Not Found</h1>
        <p>The page you are looking for does not exist.</p>
      </div>
    </div>
  );
};

export default NotFound;
