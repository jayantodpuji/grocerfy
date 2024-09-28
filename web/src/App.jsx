import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Routes, useNavigate } from 'react-router-dom';
import LoginForm from './LoginForm';
import RegisterForm from './RegisterForm';
import Dashboard from './components/dashboard';
import { isAuthenticated } from './utilities/auth';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<AuthPage />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </Router>
  );
}

const AuthPage = () => {
  const navigate = useNavigate();

  useEffect(() => {
    if (isAuthenticated()) {
      navigate('/dashboard');
    }
  }, [navigate])

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
};

export default App;
