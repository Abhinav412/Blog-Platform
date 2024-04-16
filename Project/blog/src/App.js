import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Welcome from "./welcome";
import Home from "./home";
import CreatePost from "./CreatePost";
import AllPosts from "./Allposts";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Welcome/>}/>
        <Route path="/home" element={<Home/>}/>
        <Route path="/createpost" element={<CreatePost/>}/>
        <Route path="/allposts" element={<AllPosts/>}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
