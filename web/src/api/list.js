const API_BASE_URL = 'http://localhost:3000/api/v1/secured';

const handleResponse = async (response) => {
  if (response.status === 401) {
    throw new Error('Unauthorized');
  }

  if (!response.ok) {
    throw new Error(`API request failed with status ${response.status}`);
  }

  const responseText = await response.text();

  if (!responseText) {
    console.log('Empty response received');
    return null;
  }

  try {
    const responseData = JSON.parse(responseText);
    if (!responseData.success) {
      throw new Error(responseData.message || 'Operation failed');
    }
    return responseData;
  } catch (error) {
    console.error('Error parsing JSON:', error);
    throw new Error('Invalid response from server');
  }
};

export const fetchLists = async () => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE_URL}/lists/`, {
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });

  const responseData = await handleResponse(response);
  return responseData?.data || [];
};

export const createNewList = async (name, description) => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE_URL}/lists/`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ name, description }),
  });

  const responseData = await handleResponse(response);
  if (response.status === 201) {
    return { name, description };
  } else {
    throw new Error('Unexpected response from server');
  }
};

export const fetchListDetails = async (id) => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE_URL}/lists/${id}`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
  });

  const responseData = await handleResponse(response);
  return responseData?.data;
};

export const createNewItem = async (listId, name, quantity, unit) => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE_URL}/items/`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ listId, name, quantity, unit }),
  });

  const responseData = await handleResponse(response);
  if (response.ok) {
    return responseData?.data;
  } else {
    throw new Error('Failed to add item');
  }
};

export const deleteList = async (listId) => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE_URL}/lists/${listId}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });

  if (!response.ok) {
    throw new Error(`Failed to delete list. Status: ${response.status}`);
  }
};

export const toggleIsPurchased = async (itemId) => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE_URL}/items/${itemId}/toggle`, {
    method: 'PATCH',
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });

  if (!response.ok) {
    throw new Error(`Failed to toggle item status. Status: ${response.status}`);
  }
};
