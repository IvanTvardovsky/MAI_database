import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
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
      <Link to="/main" style={{ textDecoration: 'none' }}>
         <button style={{ padding: '10px', margin: '10px', cursor: 'pointer', backgroundColor: '#4caf50', color: 'white', border: 'none', borderRadius: '5px' }}>
            Перейти к выбору направления
         </button>
      </Link>
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
            <Route path="/changingRating" element={<ChangingRating />} />
         </Routes>
      </BrowserRouter>
   );
};

export default App;
