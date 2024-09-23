export const Register = async (name, email, password) => {
  try {
    const response = await fetch('http://localhost:3000/api/v1/users/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, email, password }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Registration failed');
    }

    return true;
  } catch (err) {
    throw err;
  }
};
