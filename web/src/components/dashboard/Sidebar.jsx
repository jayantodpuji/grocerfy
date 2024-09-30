import React from 'react';

const Sidebar = ({ items, onItemClick, onItemRemove, onLogout }) => {
  return (
    <div className="drawer-side">
      <label htmlFor="my-drawer-2" aria-label="close sidebar" className="drawer-overlay"></label>
      <div className="flex flex-col h-full bg-base-200 text-base-content w-80">
        <div className="flex-grow overflow-y-auto">
          <ul className="menu p-2">
            <li className="mb-2">
              <h3 className="text-xl font-bold text-base-content/70 px-4 py-2">Grocery List</h3>
            </li>
            {items.map((item) => (
              <li key={item.id} className="relative group mb-2">
                <button
                  className="btn btn-block justify-between hover:bg-base-300 transition-colors duration-200"
                  onClick={() => onItemClick(item)}
                >
                  <span className="flex-grow text-left">{item.name}</span>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-6 w-6 flex-shrink-0 opacity-0 group-hover:opacity-100 transition-opacity duration-200 hover:text-red-500"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    onClick={(e) => {
                      e.stopPropagation();
                      onItemRemove(item.id);
                    }}
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </li>
            ))}
          </ul>
        </div>
        <div className="p-4 border-t border-base-300">
          <button
            className="btn btn-block btn-outline btn-error"
            onClick={onLogout}
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
