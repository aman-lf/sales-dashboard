import { useEffect, useState } from 'react';
import { Outlet } from 'react-router-dom';

import Toast from '@/components/Toast';
import Sidebar from '@/components/Sidebar';

import { api } from '@/constants/api';

import './layout.scss';

const AppLayout = () => {
  const showSidebar = true;
  const [toastProps, setToastProps] = useState({
    message: '',
    open: false,
  });

  useEffect(() => {
    const sse = new EventSource(`${import.meta.env.VITE_API_URL}${api.NOTIFICATON}`);
    sse.onmessage = (e) => {
      const message = JSON.parse(e.data).message;
      if (message) showToast(JSON.parse(e.data).message);
    };
    sse.onerror = () => {
      console.log('sse error');
      sse.close();
    };
    return () => {
      sse.close();
    };
  }, []);

  const showToast = (message, autoCloseDuration = 5000) => {
    setToastProps({
      message,
      open: true,
    });

    setTimeout(() => {
      setToastProps({ open: false, message: '' });
    }, autoCloseDuration);
  };

  return (
    <div className='main'>
      <Toast open={toastProps.open} message={toastProps.message} />

      {showSidebar ? <Sidebar /> : null}
      <Outlet />
    </div>
  );
};

export default AppLayout;
