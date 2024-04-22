import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Welcome from "./welcome";
import Home from "./home";
import CreatePostForm from "./CreatePost";
import AllPostsrender from "./Allposts";


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Welcome/>}/>
        <Route path="/home" element={<Home/>}/>
        <Route path="/createpost" element={<CreatePostForm/>}/>
        <Route path="/allposts" element={<AllPostsrender />} />

      </Routes>
    </BrowserRouter>
  );
}

export default App;
