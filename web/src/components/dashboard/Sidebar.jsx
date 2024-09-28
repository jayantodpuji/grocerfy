import React from 'react';

const Sidebar = ({ items, onItemClick }) => {
  return (
    <div className="drawer-side">
      <label htmlFor="my-drawer-2" aria-label="close sidebar" className="drawer-overlay"></label>
      <ul className="menu p-4 w-80 min-h-full bg-base-200 text-base-content">
        <li className="mb-4">
          <h3 className="text-xl font-bold">Menu</h3>
        </li>
        <li className="mb-2">
          <h3 className="text-lg font-semibold">List</h3>
        </li>
        {items.map((item) => (
          <li key={item.id}>
            <a onClick={() => onItemClick(item)} className="py-2">{item.name}</a>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Sidebar;
