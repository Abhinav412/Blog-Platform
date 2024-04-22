import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Signup() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/api/users/signup', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json,charset=utf-8', },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        throw new Error(`Signup failed with status: ${response.status}`);
      }

      // Handle successful signup (redirect to login or create post page)
      setError(null); // Clear any previous errors
      alert('Signup successful! Please login.'); // Simple alert for demo
      navigate('/login'); // Redirect to login page
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gray-100">
      {error && <p className="text-red-500 text-center">{error}</p>}  {/* Display error message if any */}
      <form onSubmit={handleSubmit} className="flex flex-col space-y-4">
        <div className="w-full">
          <label htmlFor="username" className="text-sm font-medium text-gray-700">Username</label>
          <input
            type="text"
            id="username"
            className="w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-1 focus:ring-blue-500"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div className="w-full">
          <label htmlFor="password" className="text-sm font-medium text-gray-700">Password</label>
          <input
            type="password"
            id="password"
            className="w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-1 focus:ring-blue-500"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button type="submit" className="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
          Sign Up
        </button>
      </form>
    </div>
  );
}

export default Signup;