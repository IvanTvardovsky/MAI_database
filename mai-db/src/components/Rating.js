import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import './../css/Rating.css';

const Rating = () => {
   const [placesData, setPlacesData] = useState(null);

   useEffect(() => {
      const fetchData = async () => {
         try {
            const response = await fetch('http://localhost:5050/places');
            if (!response.ok) {
               throw new Error('Ошибка при получении данных');
            }

            const data = await response.json();
            setPlacesData(data);
         } catch (error) {
            console.error('Произошла ошибка:', error.message);
         }
      };

      fetchData();
   }, []);

   const downloadReport = async () => {
      try {
         const response = await fetch('http://localhost:5050/reportPlaces');
         if (!response.ok) {
            throw new Error('Ошибка при загрузке отчета');
         }

         const blob = await response.blob();
         const url = window.URL.createObjectURL(blob);
         const a = document.createElement('a');
         a.href = url;
         a.download = 'Отчёт по рейтингу.pdf';
         document.body.appendChild(a);
         a.click();
         document.body.removeChild(a);
      } catch (error) {
         console.error('Произошла ошибка при загрузке отчета:', error.message);
      }
   };

   return (
      <div className="rating-container">
         <h2>Рейтинг вузов</h2>
         <Link to="/main" className="link-to-main">Вернуться к направлениям</Link>
         <button onClick={downloadReport} className="button-download">Скачать отчет</button>
         {placesData ? (
            <ul className="rating-list">
               {placesData.map((place, index) => (
                  <li key={index} className="rating-item">
                     <strong>{place.UniversityName}</strong>
                     <p>Место: {place.Place}</p>
                  </li>
               ))}
            </ul>
         ) : (
            <p>Loading...</p>
         )}
      </div>
   );
};

export default Rating;