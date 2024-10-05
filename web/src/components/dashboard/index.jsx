import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { isAuthenticated } from '../../utilities/auth';
import { fetchLists, createNewList } from '../../api/list';
import Sidebar from './Sidebar';
import Detail from './Detail';
import NewList from './NewList';

const Dashboard = () => {
  const navigate = useNavigate();
  const [items, setItems] = useState([]);
  const [selectedItem, setSelectedItem] = useState(null);
  const [isCreatingNewList, setIsCreatingNewList] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!isAuthenticated()) {
      navigateToHome();
    } else {
      loadLists();
    }
  }, []);

  const navigateToHome = () => {
    localStorage.removeItem('token');
    navigate('/');
  };

  const loadLists = async () => {
    try {
      const data = await fetchLists();
      setItems(data);
      setIsLoading(false);
    } catch (err) {
      handleError(err);
    }
  };

  const handleError = (err) => {
    console.error('Error occurred:', err);
    if (err.message === 'Unauthorized') {
      navigateToHome();
    } else {
      setError(err.message);
      setIsLoading(false);
    }
  };

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

  const handleSaveNewList = async (name, description) => {
    try {
      setIsLoading(true);
      await createNewList(name, description);
      await loadLists();
      setIsCreatingNewList(false);
    } catch (err) {
      handleError(err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleCancelNewList = () => {
    setIsCreatingNewList(false);
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

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
