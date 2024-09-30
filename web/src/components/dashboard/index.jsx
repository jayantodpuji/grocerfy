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
    { id: 5, name: 'Item 5', description: 'This is the description for Item 5' },
    { id: 6, name: 'Item 6', description: 'This is the description for Item 6' },
    { id: 7, name: 'Item 7', description: 'This is the description for Item 7' },
    { id: 8, name: 'Item 8', description: 'This is the description for Item 8' },
    { id: 9, name: 'Item 9', description: 'This is the description for Item 9' },
    { id: 10, name: 'Item 10', description: 'This is the description for Item 10' },
    { id: 11, name: 'Item 11', description: 'This is the description for Item 11' },
    { id: 12, name: 'Item 12', description: 'This is the description for Item 12' },
    { id: 13, name: 'Item 13', description: 'This is the description for Item 13' },
    { id: 14, name: 'Item 14', description: 'This is the description for Item 14' },
    { id: 15, name: 'Item 15', description: 'This is the description for Item 15' },
    { id: 16, name: 'Item 16', description: 'This is the description for Item 16' },
    { id: 17, name: 'Item 17', description: 'This is the description for Item 17' },
    { id: 18, name: 'Item 18', description: 'This is the description for Item 18' },
    { id: 19, name: 'Item 19', description: 'This is the description for Item 19' },
    { id: 20, name: 'Item 20', description: 'This is the description for Item 20' },
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
