import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import './../css/Rating.css'

const ChangingRating = () => {
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

   return (
      <div className="rating-container">
         <h2>Рейтинг вузов</h2>
         <Link to="/main">Вернуться к направлениям</Link>
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

export default ChangingRating;