import React, { useState } from 'react';
import axios from 'axios';

const UniversityForm = () => {
   const [passingScore, setPassingScore] = useState('');
   const [budgetPlaces, setBudgetPlaces] = useState('');
   const [backendResponse, setBackendResponse] = useState('');

   const handleSubmit = async (e) => {
      e.preventDefault();

      const requestData = {
         passingScore: parseInt(passingScore),
         budgetPlaces: parseInt(budgetPlaces),
      };

      try {
         const response = await axios.get('http://localhost:5050/test', requestData);
         setBackendResponse(response.data);
      } catch (error) {
         console.error('Ошибка при отправке запроса:', error);
      }
   };

   return (
      <div>
         <form onSubmit={handleSubmit}>
            <label>
               Проходной балл:
               <input type="text" value={passingScore} onChange={(e) => setPassingScore(e.target.value)} />
            </label>
            <br />
            <label>
               Бюджетных мест:
               <input type="text" value={budgetPlaces} onChange={(e) => setBudgetPlaces(e.target.value)} />
            </label>
            <br />
            <button type="submit">Отправить запрос</button>
         </form>

         {backendResponse && (
            <div>
               <h3>Ответ от попы:</h3>
               <p>{backendResponse}</p>
            </div>
         )}
      </div>
   );
};

export default UniversityForm;
