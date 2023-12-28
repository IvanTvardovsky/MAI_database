import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Navbar from './components/NavBar';
import "./css/App.css"
import Main from "./components/FilterPanel";
import AdminPanel from "./components/AdminPanel";
import PlaceHolder from "./components/placeHolder";
import Rating from "./components/Rating";
import ChangingRating from "./components/ChangingRating";

const Home = () => (
   <div>
      <h4 className="home">Куда поступать? Сайт для поступающих о вузах и поступлении на программы бакалавриата и специалитета.</h4>
   </div>
);

const App = () => {
   return (
      <BrowserRouter>
         <Navbar />
         <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/main" element={<Main />} />
            <Route path="/admin" element={<AdminPanel />} />
            <Route path="/rating" element={<Rating />} />
            <Route path="/changeRating" element={<ChangingRating />} />
         </Routes>
      </BrowserRouter>
   );
};

export default App;
