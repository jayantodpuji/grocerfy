import React, { useState } from 'react';

const Detail = ({ selectedItem }) => {
  const [newItemText, setNewItemText] = useState('');
  const [checklist, setChecklist] = useState([
    { text: 'First dummy item', checked: false },
    { text: 'Second dummy item', checked: true },
    { text: 'Third dummy item', checked: false },
  ]);

  const handleCheckboxChange = (index) => {
    const updatedChecklist = [...checklist];
    updatedChecklist[index].checked = !updatedChecklist[index].checked;
    setChecklist(updatedChecklist);
  };

  const handleAddItem = () => {
    if (newItemText.trim()) {
      setChecklist([...checklist, { text: newItemText.trim(), checked: false }]);
      setNewItemText('');
    }
  };

  if (!selectedItem) {
    return (
      <div className="pt-16 lg:pt-8 w-full max-w-2xl mx-auto">
        <p className="text-center text-gray-500">Select an item from the menu to view details</p>
      </div>
    );
  }

  return (
    <div className="pt-16 lg:pt-8 w-full max-w-2xl mx-auto">
      <h2 className="text-2xl font-bold mb-4">{selectedItem.name}</h2>

      <ul className="mb-4">
        {checklist.map((item, index) => (
          <li key={index} className="flex items-center mb-2">
            <input
              type="checkbox"
              checked={item.checked}
              onChange={() => handleCheckboxChange(index)}
              className="mr-2"
            />
            <div className="collapse bg-white-base-100 collapse-arrow">
              <input type="checkbox" />
              <div className="collapse-title text-m font-medium">Click me to show/hide content</div>
              <div className="collapse-content">
                <p>hello</p>
              </div>
            </div>
          </li>
        ))}
      </ul>

      <div className="flex">
        <button
          onClick={handleAddItem}
          className="bg-blue-500 text-white px-4 py-2 rounded-r hover:bg-blue-600"
        >
          Add Item
        </button>
      </div>
    </div>
  );
};

export default Detail;
