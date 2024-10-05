import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { isAuthenticated } from '../../utilities/auth';
import Sidebar from './Sidebar';
import Detail from './Detail';
import NewList from './NewList';

const Dashboard = () => {
  const navigate = useNavigate();
  const [items, setItems] = useState([
    { id: 1, name: 'Item 1', description: 'This is the description for Item 1' },
    // ... (other items)
  ]);

  const [selectedItem, setSelectedItem] = useState(null);
  const [isCreatingNewList, setIsCreatingNewList] = useState(false);

  useEffect(() => {
    if (!isAuthenticated()) {
      navigate('/');
    }
  }, [navigate]);

  const handleItemClick = (item) => {
    setSelectedItem(item);
    setIsCreatingNewList(false);
  };

  const handleItemRemove = (itemId) => {
    setItems(items.filter(item => item.id !== itemId));
    if (selectedItem && selectedItem.id === itemId) {
      setSelectedItem(null);
    }
  };

  const handleCreateNewList = () => {
    setIsCreatingNewList(true);
    setSelectedItem(null);
  };

  const handleSaveNewList = (name, description) => {
    const newId = items.length > 0 ? Math.max(...items.map(item => item.id)) + 1 : 1;
    const newItem = { id: newId, name, description };
    setItems([...items, newItem]);
    setSelectedItem(newItem);
    setIsCreatingNewList(false);
  };

  const handleCancelNewList = () => {
    setIsCreatingNewList(false);
  };

  return (
    <div className="drawer lg:drawer-open">
      <input id="my-drawer-2" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content flex flex-col">
        <label htmlFor="my-drawer-2" className="btn btn-primary drawer-button lg:hidden mb-4">Open menu</label>
        {isCreatingNewList ? (
          <NewList onSave={handleSaveNewList} onCancel={handleCancelNewList} />
        ) : (
          <Detail selectedItem={selectedItem} />
        )}
      </div>
      <Sidebar
        items={items}
        onItemClick={handleItemClick}
        onItemRemove={handleItemRemove}
        onCreateNewList={handleCreateNewList}
      />
    </div>
  );
};

export default Dashboard;
