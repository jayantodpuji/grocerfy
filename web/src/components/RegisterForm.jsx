import React, { useState } from 'react';
import { Register } from '../api/register';

const RegisterForm = ({ toggleForm }) => {
  const [fullName, setFullName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (password !== confirmPassword) {
      setError("Passwords do not match");
      return;
    }

    try {
      const success = await Register(fullName, email, password);

      if (success) {
        setFullName('');
        setEmail('');
        setPassword('');
        setConfirmPassword('');
        setError(null);
        toggleForm();
      }
    } catch (err) {
      setError(err.message);
    }
  }

  return (
    <div>
      <h2 className="text-4xl font-bold mb-4">Sign Up</h2>
        <form onSubmit={handleSubmit}>
        <div className="mb-4">
            <label htmlFor="fullName" className="block text-sm font-medium text-gray-700">Full Name</label>
            <input
              type="text"
              id="fullName"
              value={fullName}
              onChange={(e) => setFullName(e.target.value)}
              className="mt-1 block w-full p-3 border border-gray-300 rounded-md"
              placeholder="John Doe"
              required
            />
          </div>
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
          <div className="mb-4">
            <label htmlFor="confirm-password" className="block text-sm font-medium text-gray-700">Confirm Password</label>
            <input
              type="password"
              id="confirm-password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              className="mt-1 block w-full p-3 border border-gray-300 rounded-md"
              placeholder="Confirm your password"
              required
            />
          </div>
          {error && <p className="text-red-500 mb-4">{error}</p>}
          <button
            type="submit"
            className="w-full py-3 px-4 bg-green-600 text-white font-semibold rounded-md hover:bg-green-700"
          >
            Sign Up
          </button>
        </form>
      <p className="mt-4 text-sm text-center">
        Already have an account?{' '}
        <button onClick={toggleForm} className="text-blue-600 hover:text-blue-800 font-semibold">
          Login
        </button>
      </p>
    </div>
  );
};

export default RegisterForm;
