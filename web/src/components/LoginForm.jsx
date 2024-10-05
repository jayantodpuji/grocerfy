import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Login } from '../api/login';

const LoginForm = ({ toggleForm }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const { token } = await Login(email, password);

      localStorage.setItem('token', token);

      navigate('/dashboard');
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="w-full max-w-sm">
      <h2 className="text-4xl font-bold mb-4">Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="mt-1 block w-full p-3 border border-gray-300 rounded-md"
            placeholder="you@example.com"
            required
          />
        </div>
        <div className="mb-4">
          <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="mt-1 block w-full p-3 border border-gray-300 rounded-md"
            placeholder="Enter your password"
            required
          />
        </div>
        {error && <p className="text-red-500 mb-4">{error}</p>}
        <button
          type="submit"
          className="w-full py-3 px-4 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700"
        >
          Login
        </button>
      </form>
      <p className="mt-4 text-sm text-center">
        Don't have an account?{' '}
        <button onClick={toggleForm} className="text-blue-600 hover:text-blue-800 font-semibold">
          Sign Up
        </button>
      </p>
    </div>
  );
};

export default LoginForm;
