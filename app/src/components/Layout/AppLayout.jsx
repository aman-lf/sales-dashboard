import React, { useContext, useState } from 'react';
import { Outlet, useOutletContext } from 'react-router-dom';

import Sidebar from '../Sidebar/Sidebar';

import './layout.scss';

const AppLayout = () => {
  const [showSidebar, setShowSidebar] = useState(true);
  const [toastProps, setToastProps] = useState({
    message: '',
    open: false,
  });

  const hideSidebar = () => {
    setShowSidebar(false);
  };

  const showToast = (message, autoCloseDuration = 3000) => {
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
      {/* <Toast open={toastProps.open} message={toastProps.message} /> */}

      {showSidebar ? <Sidebar /> : null}
      <Outlet context={{ hideSidebar, showToast }} />
    </div>
  );
};

export function useLayoutContext() {
  return useOutletContext();
}

export default AppLayout;
