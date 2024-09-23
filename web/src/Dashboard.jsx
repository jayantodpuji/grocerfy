import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { isAuthenticated } from './utilities/auth';

const Dashboard = () => {
  const navigate = useNavigate();

  useEffect(() => {
    if (!isAuthenticated()) {
      navigate('/');
    }
  }, [navigate]);

  return (
    <div className="h-screen flex items-center justify-center">
      <h1 className="text-4xl font-bold">Welcome to the Dashboard!</h1>
    </div>
  );
};

export default Dashboard;
