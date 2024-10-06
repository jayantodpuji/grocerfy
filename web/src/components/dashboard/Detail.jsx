import React, { useState } from 'react';
import EmptyState from '../EmptyState';
import { createNewItem, toggleIsPurchased } from '../../api/list';

const Detail = ({ selectedItem, onRefresh }) => {
  const [showForm, setShowForm] = useState(false);
  const [newItem, setNewItem] = useState({ name: '', quantity: '', unit: '' });
  const [items, setItems] = useState(selectedItem?.items || []);

  if (!selectedItem) {
    return (
      <div className="pt-16 lg:pt-8 w-full max-w-2xl mx-auto">
        <EmptyState/>
      </div>
    );
  }

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setNewItem({ ...newItem, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await createNewItem(
        selectedItem.id,
        newItem.name,
        Number(newItem.quantity),
        newItem.unit
      );

      setNewItem({ name: '', quantity: '', unit: '' });
      setShowForm(false);
      onRefresh();
    } catch (error) {
      console.error('Error adding item:', error);
    }
  };

  const handleToggleIsPurchased = async (itemId) => {
    try {
      await toggleIsPurchased(itemId);
      setItems(items.map(item =>
        item.id === itemId ? { ...item, isPurchased: !item.isPurchased } : item
      ));
    } catch (error) {
      console.error('Error toggling item status:', error);
    }
  };

  return (
    <div className="pt-16 lg:pt-8 w-full max-w-2xl mx-auto">
      <h2 className="text-2xl font-bold mb-2">{selectedItem.name}</h2>
      <p className="text-gray-600 mb-4">{selectedItem.description}</p>

      <ul className="space-y-2 mb-4">
        {items.map((item) => (
          <li key={item.id} className="flex items-center space-x-2">
            <label className="label cursor-pointer">
              <input
                type="checkbox"
                checked={item.isPurchased}
                className="checkbox checkbox-primary"
                onChange={() => handleToggleIsPurchased(item.id)}
              />
              <span className="label-text ml-2">
                {item.name} - {item.quantity} {item.unit}
              </span>
            </label>
          </li>
        ))}
      </ul>

      {!showForm && (
        <button
          className="btn btn-primary"
          onClick={() => setShowForm(true)}
        >
          Add New Item
        </button>
      )}

      {showForm && (
        <form onSubmit={handleSubmit} className="form-control w-full max-w-xs">
          <label className="label">
            <span className="label-text">Name</span>
          </label>
          <input
            type="text"
            name="name"
            value={newItem.name}
            onChange={handleInputChange}
            className="input input-bordered w-full max-w-xs"
            required
          />

          <label className="label">
            <span className="label-text">Quantity</span>
          </label>
          <input
            type="number"
            name="quantity"
            value={newItem.quantity}
            onChange={handleInputChange}
            className="input input-bordered w-full max-w-xs"
            required
          />

          <label className="label">
            <span className="label-text">Unit</span>
          </label>
          <input
            type="text"
            name="unit"
            value={newItem.unit}
            onChange={handleInputChange}
            className="input input-bordered w-full max-w-xs"
            required
          />

          <div className="mt-4">
            <button type="submit" className="btn btn-primary mr-2">Save</button>
            <button type="button" className="btn btn-ghost" onClick={() => setShowForm(false)}>Cancel</button>
          </div>
        </form>
      )}
    </div>
  );
};

export default Detail;
