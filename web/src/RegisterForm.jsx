import React from 'react';

const RegisterForm = ({ toggleForm }) => {
  return (
    <div>
      <h2 className="text-4xl font-bold mb-4">Sign Up</h2>
      <form>
        <div className="mb-4">
          <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email</label>
          <input type="email" id="email" className="mt-1 block w-full p-3 border border-gray-300 rounded-md" placeholder="you@example.com" />
        </div>
        <div className="mb-4">
          <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
          <input type="password" id="password" className="mt-1 block w-full p-3 border border-gray-300 rounded-md" placeholder="Enter your password" />
        </div>
        <div className="mb-4">
          <label htmlFor="confirm-password" className="block text-sm font-medium text-gray-700">Confirm Password</label>
          <input type="password" id="confirm-password" className="mt-1 block w-full p-3 border border-gray-300 rounded-md" placeholder="Confirm your password" />
        </div>
        <button type="submit" className="w-full py-3 px-4 bg-green-600 text-white font-semibold rounded-md hover:bg-green-700">Sign Up</button>
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
