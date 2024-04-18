import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Welcome from "./welcome";
import Home from "./home";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Welcome/>}/>
        <Route path="/home" element={<Home/>}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
