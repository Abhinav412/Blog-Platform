import React from 'react';

const AuthContext = React.createContext({
  loggedIn: false,
  token: null,
  // Add other auth-related state variables here
});

export default AuthContext;