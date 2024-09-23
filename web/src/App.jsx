import React, { useState } from 'react';
import LoginForm from './LoginForm';
import RegisterForm from './RegisterForm';

function App() {
  const [isLogin, setIsLogin] = useState(true);

  const toggleForm = () => {
    setIsLogin(!isLogin);
  };

  return (
    <div className="h-screen flex">
      <div className="w-1/2 bg-gray-200 p-8 flex items-center justify-center">
        <div>
          <h1 className="text-4xl font-bold">Welcome to Grocerfy</h1>
          <p className="mt-4 text-lg">A simple Grocery Todo List Application</p>
        </div>
      </div>
      <div className="w-1/2 p-8 flex items-center justify-left">
        <div>
          {isLogin ? (
            <LoginForm toggleForm={toggleForm} />
          ) : (
            <RegisterForm toggleForm={toggleForm} />
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
