import { useNavigate } from 'react-router-dom';

import BrandIcon from '@images/brand_icon.png';
import SalesLogo from '@images/sales_icon.png';
import ProductIcon from '@images/product_icon.png';
import DashboardIcon from '@images/dashboard_icon.png';

import { PATHS } from '../../constants/routes';

import './sidebar.scss';

const Sidebar = () => {
  const navigate = useNavigate();

  return (
    <div className='sidebar'>
      <div
        className='sidebar__header'
        onClick={() => {
          navigate(PATHS.HOME_PATH);
        }}
      >
        <img src={SalesLogo} alt='sales-dashboard-logo' className='sidebar__logo' />
        <h3>Sales Dashboard</h3>
      </div>
      <div>
        <div
          className='sidebar__item'
          onClick={() => {
            navigate(PATHS.HOME_PATH);
          }}
        >
          <img src={DashboardIcon} alt='dashboard-logo' className='sidebar__logo' />
          <p>Dashboard</p>
        </div>
        <div
          className='sidebar__item'
          onClick={() => {
            navigate(PATHS.PRODUCTS_PATH);
          }}
        >
          <img src={ProductIcon} alt='product-logo' className='sidebar__logo' />
          <p>Sales By Product</p>
        </div>
        <div
          className='sidebar__item'
          onClick={() => {
            navigate(PATHS.BRAND_PATH);
          }}
        >
          <img src={BrandIcon} alt='brand-logo' className='sidebar__logo' />
          <p>Sales By Brand</p>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
