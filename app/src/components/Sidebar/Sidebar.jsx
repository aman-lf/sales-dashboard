import React from 'react';

import BrandIcon from '@images/brand_icon.png';
import SalesLogo from '@images/sales_icon.png';
import ProductIcon from '@images/product_icon.png';
import DashboardIcon from '@images/dashboard_icon.png';

import './sidebar.scss';

const Sidebar = () => {
  return (
    <div className='sidebar'>
      <div className='sidebar__header'>
        <img src={SalesLogo} alt='sales-dashboard-logo' className='sidebar__logo' />
        <h3>Sales Dashboard</h3>
      </div>
      <div>
        <div className='sidebar__item'>
          <img src={DashboardIcon} alt='dashboard-logo' className='sidebar__logo' />
          <p>Dashboard</p>
        </div>
        <div className='sidebar__item'>
          <img src={ProductIcon} alt='product-logo' className='sidebar__logo' />
          <p>Sales By Product</p>
        </div>
        <div className='sidebar__item'>
          <img src={BrandIcon} alt='brand-logo' className='sidebar__logo' />
          <p>Sales By Brand</p>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
