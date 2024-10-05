import React, { useState } from 'react';
import EmptyState from '../EmptyState';

const Detail = ({ selectedItem }) => {
  const [checklist, setChecklist] = useState([
    { id: 1, text: 'First dummy item', category: 'Category A', name: 'Item 1', quantity: 1, price: 10 },
    { id: 2, text: 'Second dummy item', category: 'Category B', name: 'Item 2', quantity: 2, price: 20 },
    { id: 3, text: 'Third dummy item', category: 'Category C', name: 'Item 3', quantity: 3, price: 30 },
  ]);
  const [editingItem, setEditingItem] = useState(null);
  const [isAddingItem, setIsAddingItem] = useState(false);
  const [newItem, setNewItem] = useState({
    text: '',
    category: '',
    name: '',
    quantity: 0,
    price: 0
  });

  const handleAddItem = () => {
    setIsAddingItem(true);
  };

  const handleSaveNewItem = () => {
    if (newItem.text.trim()) {
      setChecklist([...checklist, { ...newItem, id: Date.now() }]);
      setNewItem({
        text: '',
        category: '',
        name: '',
        quantity: 0,
        price: 0
      });
      setIsAddingItem(false);
    }
  };

  const handleCancelNewItem = () => {
    setIsAddingItem(false);
    setNewItem({
      text: '',
      category: '',
      name: '',
      quantity: 0,
      price: 0
    });
  };

  const handleUpdate = (id) => {
    setEditingItem(checklist.find(item => item.id === id));
  };

  const handleSave = () => {
    setChecklist(checklist.map(item =>
      item.id === editingItem.id ? editingItem : item
    ));
    setEditingItem(null);
    // TODO: Implement actual database update logic here
  };

  const handleCancel = () => {
    setEditingItem(null);
  };

  const handleInputChange = (e, itemType = 'editing') => {
    const { name, value } = e.target;
    if (itemType === 'editing') {
      setEditingItem({ ...editingItem, [name]: value });
    } else {
      setNewItem({ ...newItem, [name]: value });
    }
  };

  const handleRemove = (id) => {
    setChecklist(checklist.filter(item => item.id !== id));
    if (editingItem && editingItem.id === id) {
      setEditingItem(null);
    }
    // TODO: Implement actual database removal logic here
  };

  if (!selectedItem) {
    return (
      <div className="pt-16 lg:pt-8 w-full max-w-2xl mx-auto">
        <EmptyState/>
      </div>
    );
  }

  const renderForm = (item, itemType = 'editing') => (
    <form className="space-y-2 text-sm">
      <div className="form-control">
        <label className="label">
          <span className="label-text text-xs">Text</span>
        </label>
        <input
          type="text"
          name="text"
          value={item.text}
          onChange={(e) => handleInputChange(e, itemType)}
          placeholder="Item text"
          className="input input-bordered input-sm w-full"
        />
      </div>
      <div className="form-control">
        <label className="label">
          <span className="label-text text-xs">Category</span>
        </label>
        <select
          name="category"
          value={item.category}
          onChange={(e) => handleInputChange(e, itemType)}
          className="select select-bordered select-sm w-full"
        >
          <option disabled value="">Select a category</option>
          <option value="Category A">Category A</option>
          <option value="Category B">Category B</option>
          <option value="Category C">Category C</option>
        </select>
      </div>
      <div className="form-control">
        <label className="label">
          <span className="label-text text-xs">Name</span>
        </label>
        <input
          type="text"
          name="name"
          value={item.name}
          onChange={(e) => handleInputChange(e, itemType)}
          placeholder="Name"
          className="input input-bordered input-sm w-full"
        />
      </div>
      <div className="form-control">
        <label className="label">
          <span className="label-text text-xs">Quantity</span>
        </label>
        <input
          type="number"
          name="quantity"
          value={item.quantity}
          onChange={(e) => handleInputChange(e, itemType)}
          placeholder="Quantity"
          className="input input-bordered input-sm w-full"
        />
      </div>
      <div className="form-control">
        <label className="label">
          <span className="label-text text-xs">Price</span>
        </label>
        <input
          type="number"
          name="price"
          value={item.price}
          onChange={(e) => handleInputChange(e, itemType)}
          placeholder="Price"
          className="input input-bordered input-sm w-full"
        />
      </div>
      <div className="flex justify-end space-x-2 mt-4">
        <button
          type="button"
          onClick={itemType === 'editing' ? handleSave : handleSaveNewItem}
          className="btn btn-primary btn-xs"
        >
          Save
        </button>
        <button
          type="button"
          onClick={itemType === 'editing' ? handleCancel : handleCancelNewItem}
          className="btn btn-ghost btn-xs"
        >
          Cancel
        </button>
      </div>
    </form>
  );

  return (
    <div className="pt-16 lg:pt-8 w-full max-w-2xl mx-auto">
      <h2 className="text-xl font-bold mb-4">{selectedItem.name}</h2>

      <button
        onClick={handleAddItem}
        className="btn btn-primary btn-sm mb-4"
      >
        Add Item
      </button>

      {isAddingItem && (
        <div className="card bg-base-100 shadow-sm mb-4">
          <div className="card-body p-4">
            <h3 className="card-title text-sm mb-2">New Item</h3>
            {renderForm(newItem, 'new')}
          </div>
        </div>
      )}

      <ul className="space-y-2">
        {checklist.map((item) => (
          <li key={item.id} className="card bg-base-100 shadow-sm">
            <div className="card-body p-2">
              <div className="collapse collapse-arrow">
                <input type="checkbox" />
                <div className="collapse-title text-sm font-medium py-2 pr-10 pl-0 min-h-0">
                  {item.text}
                </div>
                <div className="collapse-content text-xs px-0">
                  {editingItem && editingItem.id === item.id ? (
                    renderForm(editingItem)
                  ) : (
                    <div>
                      <p><strong>Category:</strong> {item.category}</p>
                      <p><strong>Name:</strong> {item.name}</p>
                      <p><strong>Quantity:</strong> {item.quantity}</p>
                      <p><strong>Price:</strong> ${item.price}</p>
                      <div className="flex justify-end space-x-2 mt-2">
                        <button
                          onClick={() => handleUpdate(item.id)}
                          className="btn btn-primary btn-xs"
                        >
                          Update
                        </button>
                        <button
                          onClick={() => handleRemove(item.id)}
                          className="btn btn-error btn-xs"
                        >
                          Remove
                        </button>
                      </div>
                    </div>
                  )}
                </div>
              </div>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Detail;
