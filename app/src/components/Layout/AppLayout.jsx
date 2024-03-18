import React, { useEffect, useState } from 'react';
import { Outlet } from 'react-router-dom';

import Toast from '../Toast/Toast';
import Sidebar from '../Sidebar/Sidebar';

import './layout.scss';

const AppLayout = () => {
  const [showSidebar, setShowSidebar] = useState(true);
  const [toastProps, setToastProps] = useState({
    message: '',
    open: false,
  });

  useEffect(() => {
    const sse = new EventSource(`${import.meta.env.VITE_API_URL}/new-file-notification`);
    sse.onmessage = (e) => {
      const message = JSON.parse(e.data).message;
      if (message) showToast(JSON.parse(e.data).message);
    };
    sse.onerror = () => {
      // error log here
      console.log('sse error');
      sse.close();
    };
    return () => {
      sse.close();
    };
  }, []);

  const hideSidebar = () => {
    setShowSidebar(false);
  };

  const showToast = (message, autoCloseDuration = 5000) => {
    setToastProps({
      message,
      open: true,
    });

    const timeout = setTimeout(() => {
      setToastProps({ open: false, message: '' });
    }, autoCloseDuration);
  };

  return (
    <div className='main'>
      <Toast open={toastProps.open} message={toastProps.message} />

      {showSidebar ? <Sidebar /> : null}
      <Outlet context={{ hideSidebar, showToast }} />
    </div>
  );
};

export default AppLayout;
