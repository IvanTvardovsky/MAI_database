import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import './../css/Rating.css';
import "./../css/ChangingRating.css"

const ChangingRating = () => {
   const [placesData, setPlacesData] = useState(null);
   const [isModalOpen, setModalOpen] = useState(false);
   const [newPlace, setNewPlace] = useState('');
   const [selectedUniversity, setSelectedUniversity] = useState(null);

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

   const handleOpenModal = (university) => {
      setSelectedUniversity(university);
      setModalOpen(true);
   };

   const handleCloseModal = () => {
      setSelectedUniversity(null);
      setModalOpen(false);
      setNewPlace('');
   };

   const handleInputChange = (event) => {
      setNewPlace(event.target.value);
   };

   const handleSendData = async () => {
      try {
         const response = await fetch('http://localhost:5050/updatePlace', {
            method: 'PUT',
            headers: {
               'Content-Type': 'application/json',
            },
            body: JSON.stringify({
               UniversityName: selectedUniversity,
               Place: parseInt(newPlace),
            }),
         });

         if (!response.ok) {
            throw new Error('Ошибка при отправке данных');
         }

         // Обновление данных после успешной отправки
         const updatedPlaces = [...placesData];
         const index = updatedPlaces.findIndex((place) => place.UniversityName === selectedUniversity);
         updatedPlaces[index].Place = parseInt(newPlace);
         setPlacesData(updatedPlaces);

         // Закрытие модального окна
         handleCloseModal();
      } catch (error) {
         console.error('Произошла ошибка:', error.message);
      }
   };

   return (
      <div className="rating-container">
         <h2>Рейтинг вузов</h2>
         <Link to="/admin">Вернуться к админ-панели</Link>

         {/* Модальное окно для ввода нового места */}
         {isModalOpen && (
            <div className="modal">
               <div className="modal-content">
               <span className="close" onClick={handleCloseModal}>
                  &times;
               </span>
                  <h3>Изменить место для {selectedUniversity}</h3>
                  <label>
                     Новое место:
                     <input type="number" value={newPlace} onChange={handleInputChange} />
                  </label>
                  <button onClick={handleSendData}>Отправить</button>
               </div>
            </div>
         )}

         {placesData ? (
            <ul className="rating-list">
               {placesData.map((place, index) => (
                  <li key={index} className={`rating-item ${place.UniversityName === selectedUniversity ? 'changed' : ''}`}>
                     <strong>{place.UniversityName}</strong>
                     <p>Место: {place.Place}</p>
                     <button onClick={() => handleOpenModal(place.UniversityName)}>Изменить</button>
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
