import React from 'react';
import "./../css/App.css"

const UniversityListAdmin = ({ universities, onEditClick }) => {
   return (
      <div className="university-list">
         {universities.map((university) => (
            <div key={university.id} className="university-item">
               <button onClick={() => onEditClick(university.id)}>Изменить</button>
               <h3>{university.name}</h3>
               <div className="info-block">
                  <div className="info-row">
                     <p className="info-label">Цена за обучение:</p>
                     <p className="info-value">{university.price}</p>
                  </div>
                  <div className="info-row">
                     <p className="info-label">Количество мест:</p>
                     <p className="info-value">{university.places}</p>
                  </div>
                  <div className="info-row">
                     <p className="info-label">Проходной балл:</p>
                     <p className="info-value">{university.passingScore}</p>
                  </div>
                  <div className="info-row">
                     <p className="info-label">Код направления:</p>
                     <p className="info-value">{university.code}</p>
                  </div>
                  <div className="info-row">
                     <p className="info-label">Военная кафедра:</p>
                     <p className="info-value">{university.university.has_military_department ? 'Присутствует' : 'Отсутствует'}</p>
                  </div>
                  <div className="info-row">
                     <p className="info-label">Государственный институт:</p>
                     <p className="info-value">{university.university.is_state_university ? 'Да' : 'Нет'}</p>
                  </div>
                  <div className="university-details">
                     <p className="info-label">Университет:</p>
                     <p className="university-name">{university.university.name}</p>
                  </div>
               </div>
               <div className="subjects">
                  <p className="info-label">Предметы:</p>
                  <ul>
                     {university.subjects.map((subject) => (
                        <li key={subject.id} className="subjectList">
                           {subject.name} {subject.is_choosable ? '(Выборочный)' : ''}
                        </li>
                     ))}
                  </ul>
               </div>
            </div>
         ))}
      </div>
   );
};

export default UniversityListAdmin;
