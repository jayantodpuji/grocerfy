import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { isAuthenticated } from '../../utilities/auth';
import Sidebar from './Sidebar';
import Detail from './Detail';

const Dashboard = () => {
  const navigate = useNavigate();
  const [items, setItems] = useState([
    { id: 1, name: 'Item 1', description: 'This is the description for Item 1' },
    { id: 2, name: 'Item 2', description: 'This is the description for Item 2' },
    { id: 3, name: 'Item 3', description: 'This is the description for Item 3' },
    { id: 4, name: 'Item 4', description: 'This is the description for Item 4' },
  ]);
  const [selectedItem, setSelectedItem] = useState(null);

  useEffect(() => {
    if (!isAuthenticated()) {
      navigate('/');
    }
  }, [navigate]);

  const handleItemClick = (item) => {
    setSelectedItem(item);
  };

  const handleItemRemove = (itemId) => {
    setItems(items.filter(item => item.id !== itemId));
    if (selectedItem && selectedItem.id === itemId) {
      setSelectedItem(null);
    }
  };

  return (
    <div className="drawer lg:drawer-open">
      <input id="my-drawer-2" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content flex flex-col items-start p-4">
        <label htmlFor="my-drawer-2" className="btn btn-primary drawer-button lg:hidden mb-4">Open menu</label>
        <Detail selectedItem={selectedItem} />
      </div>
      <Sidebar items={items} onItemClick={handleItemClick} onItemRemove={handleItemRemove} />
    </div>
  );
};

export default Dashboard;
