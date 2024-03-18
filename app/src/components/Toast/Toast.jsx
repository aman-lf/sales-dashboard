import { useEffect, useState } from 'react';

import './toast.scss';

const Toast = ({ message, open = false }) => {
  const [visible, setVisible] = useState(open);

  useEffect(() => {
    setVisible(open);
  }, [open]);

  return visible ? (
    <div className='Toast__container'>
      <div className='Toast__wrapper'>
        <div className='Toast__content'>
          <h3>{message}</h3>
        </div>
      </div>
    </div>
  ) : null;
};

export default Toast;
