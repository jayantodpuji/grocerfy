import React from 'react';

const LoginForm = ({ toggleForm }) => {
  return (
    <div>
      <h2 className="text-4xl font-bold mb-4">Login</h2>
      <form>
        <div className="mb-4">
          <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email</label>
          <input type="email" id="email" className="mt-1 block w-full p-3 border border-gray-300 rounded-md" placeholder="you@example.com" />
        </div>
        <div className="mb-4">
          <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
          <input type="password" id="password" className="mt-1 block w-full p-3 border border-gray-300 rounded-md" placeholder="Enter your password" />
        </div>
        <button type="submit" className="w-full py-3 px-4 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700">Login</button>
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
