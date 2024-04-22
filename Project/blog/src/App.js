import React, { useState } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Welcome from "./welcome";
import Home from "./home";
import Signup from './signup';
import Login from './login'; // Import login component
import CreatePost from './CreatePost';
import AllPostsrender from "./Allposts";


import AuthContext from './AuthContext'; // Import AuthContext

function App() {
  const [authState, setAuthState] = useState({ loggedIn: false, token: null });

  return (
    <AuthContext.Provider value={{ authState, setAuthState }}> {/* Wrap with Provider */}
      <BrowserRouter>
        <Routes>
        <Route path="/" element={<Welcome/>}/>
        <Route path="/home" element={<Home/>}/>
        <Route path="/createpost" element={<CreatePost/>}/>
        <Route path="/allposts" element={<AllPostsrender />} />
        <Route path="/login" element={<Login/>}/>
        <Route path="/signup" element={<Signup/>}/>
        </Routes>
      </BrowserRouter>
    </AuthContext.Provider>
  );
}

export default App;