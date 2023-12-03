import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Navbar from './components/NavBar';
import PlaceholderUniversityForm from './components/placeHolder';
import "./css/App.css"

const Home = () => (
   <div>
      <p>траляля тополя описание инструмента и кнопка для перехода к фильтру университетов</p>
   </div>
);

const App = () => {
   return (
      <BrowserRouter>
         <Navbar />
         <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/main" element={<PlaceholderUniversityForm />} />
            <Route path="/admin" element={<PlaceholderUniversityForm />} />
         </Routes>
      </BrowserRouter>
   );
};

export default App;
