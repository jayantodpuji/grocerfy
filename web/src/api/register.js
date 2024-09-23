export const Register = async (email, password) => {
  try {
    const response = await fetch('http://localhost:3000/api/v1/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
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
