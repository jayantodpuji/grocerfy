import React from 'react';

const EmptyState = () => {
  return (
    <div className="flex flex-col items-center justify-center h-full text-center p-4">
      <div className="mb-4 text-6xl text-gray-400">
        &#x1F4CB;
      </div>
      <h2 className="text-2xl font-semibold text-gray-700 mb-2">No List Selected</h2>
      <p className="text-gray-500 max-w-sm">
        Select a list from the sidebar to view its details, or create a new list to get started.
      </p>
    </div>
  );
};

export default EmptyState;
