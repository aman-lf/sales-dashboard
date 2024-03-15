import React, { useContext, useState } from "react";

import { Outlet, useOutletContext } from "react-router-dom";

type ContextType = {
  hideSidebar: Function;
  showToast: Function;
};

const AppLayout = () => {
  const [showSidebar, setShowSidebar] = useState(true);
  const [toastProps, setToastProps] = useState({
    message: "",
    open: false,
  });

  const hideSidebar = () => {
    setShowSidebar(false);
  };

  const showToast = (message: string, autoCloseDuration: number = 3000) => {
    setToastProps({
      message,
      open: true,
    });

    const timeout = setTimeout(() => {
      setToastProps({ open: false, message: "" });
    }, autoCloseDuration);
  };

  return (
    <div className='Main__container-wrapper'>
      {/* <Toast open={toastProps.open} message={toastProps.message} /> */}

      {/* {showSidebar ? (
        <Profile name={user?.name} avatar={user?.picture} />
      ) : null} */}
      <Outlet context={{ hideSidebar, showToast }} />
    </div>
  );
};

export function useLayoutContext() {
  return useOutletContext<ContextType>();
}

export default AppLayout;
