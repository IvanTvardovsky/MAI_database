import React, {useEffect, useState} from 'react';
import "./../css/App.css"
import queryString from 'query-string';
import UniversityListAdmin from './UniversityListAdmin';

const AdminPanel = () => {
   const [filters, setFilters] = useState({
      score: '',
      budgetSeats: '',
      paidSeats: '',
      universityRating: '',
      tuitionFee: '',
      egeSubjects: {
         Mathematics: false,
         Physics: false,
         Russian: false,
         Biology: false,
         English: false,
         Informatics: false,
         Literature: false,
         AEE: false,
      },
      isStateUniversity: false,
      hasMilitaryDepartment: false,
      paidScore: '',
      budgetScore: '',
   });

   const [queryStringResult, setQueryStringResult] = useState('');
   const [responseData, setResponseData] = useState(null);
   const [numberOfRecords, setNumberOfRecords] = useState('');

   const subjectNameMapping = {
      'Математика': 'Mathematics',
      'Физика': 'Physics',
      'Русский язык': 'Russian',
      'Биология': 'Biology',
      'Английский язык': 'English',
      'Информатика': 'Informatics',
      'Литература': 'Literature',
      'ДВИ': 'AEE',
   };

   const handleGetInfoClick = async () => {
      try {
         const response = await fetch(`http://localhost:5050/reportFilters?n=${numberOfRecords}`);
         const blob = await response.blob();
         const url = URL.createObjectURL(blob);
         window.open(url);
      } catch (error) {
         console.error('Error getting information:', error);
      }
   };

   const handleInputChange = (e) => {
      const {name, value} = e.target;
      setFilters((prevFilters) => ({
         ...prevFilters,
         [name]: value,
      }));
   };

   const handleCheckboxChange = (e) => {
      const { name, checked } = e.target;
      const englishSubjectName = subjectNameMapping[name] || name;

      if (englishSubjectName in filters.egeSubjects) {
         setFilters((prevFilters) => ({
            ...prevFilters,
            egeSubjects: {
               ...prevFilters.egeSubjects,
               [englishSubjectName]: checked,
            },
         }));
      } else {
         setFilters((prevFilters) => ({
            ...prevFilters,
            [englishSubjectName]: checked,
         }));
      }
   };

   const [selectedUniversity, setSelectedUniversity] = useState(null);

   const handleEditClick = (universityId) => {
      const selected = responseData.find((university) => university.id === universityId);
      setSelectedUniversity(selected);
   };

   const handleEditInputChange = (field, value) => {
      setSelectedUniversity((prev) => ({
         ...prev,
         [field]: value,
      }));
   };


   const handleSaveChanges = async () => {
      try {
         const updatedUniversity = {
            id: selectedUniversity.id,
            price: parseInt(selectedUniversity.price),
            places: parseInt(selectedUniversity.places),
            subjects: selectedUniversity.subjects.map(subject => ({ ...subject })),
            passingScore: parseInt(selectedUniversity.passingScore),
            p_type: parseInt(selectedUniversity.p_type),
            university: { ...selectedUniversity.university },
         };

         const response = await fetch(`http://localhost:5050/updateRecord`, {
            method: 'PUT',
            headers: {
               'Content-Type': 'application/json',
            },
            body: JSON.stringify(updatedUniversity),
         });

         setSelectedUniversity(null);

         handleFilterClick();
      } catch (error) {
         console.error('Error updating data:', error);
      }
   };


   const handleFilterClick = () => {
      const queryParams = {
         subjects: Object.entries(filters.egeSubjects)
            .filter(([subject, checked]) => checked && subject !== 'isStateUniversity' && subject !== 'hasMilitaryDepartment')
            .map(([subject]) => subject),
         military: filters.hasMilitaryDepartment ? 1 : undefined,
         state: filters.isStateUniversity ? 1 : undefined,
         paid_cost: filters.tuitionFee || undefined,
         paid_places: filters.paidSeats || undefined,
         budget_places: filters.budgetSeats || undefined,
         place: filters.universityRating || undefined,
         paid_score: filters.paidScore || undefined,
         budget_score: filters.budgetScore || undefined,
      };

      const nonEmptyQueryParams = Object.fromEntries(Object.entries(queryParams).filter(([_, value]) => value !== undefined));
      const generatedQueryString = queryString.stringify(nonEmptyQueryParams, {skipNull: true, arrayFormat: 'comma'});
      setQueryStringResult(generatedQueryString);

      fetchData(generatedQueryString);
   };

   const fetchData = async (query) => {
      try {
         const response = await fetch(`http://localhost:5050/getAllRecordsFilter?${query}`);
         const data = await response.json();
         setResponseData(data);
      } catch (error) {
         console.error('Error fetching data:', error);
      }
   };

   useEffect(() => {
      fetchData(queryStringResult);
   }, []);

   return (
      <div className="filter-panel">
         <div className="filter-section">

            <label className="label">
               Проходной балл на платное
               <input
                  type="number"
                  name="paidScore"
                  value={filters.paidScore}
                  onChange={handleInputChange}
               />
            </label>

            <label className="label">
               Проходной балл на бюджет
               <input
                  type="number"
                  name="budgetScore"
                  value={filters.budgetScore}
                  onChange={handleInputChange}
               />
            </label>

            <label className="label">
               Количество бюджетных мест
               <input type="number" name="budgetSeats" value={filters.budgetSeats} onChange={handleInputChange}/>
            </label>

            <label className="label">
               Количество платных мест
               <input type="number" name="paidSeats" value={filters.paidSeats} onChange={handleInputChange}/>
            </label>

            <label className="label">
               Рейтинг вуза
               <input type="number" name="universityRating" value={filters.universityRating}
                      onChange={handleInputChange}/>
            </label>

            <label className="label">
               Цена платного обучения
               <input type="number" name="tuitionFee" value={filters.tuitionFee} onChange={handleInputChange}/>
            </label>
         </div>

         <div className="filter-section">
            <label className="label">
               Сдаваемые предметы ЕГЭ
            </label>
            {Object.entries(filters.egeSubjects).map(([subject, isChecked]) => {
               if (subject !== 'isStateUniversity' && subject !== 'hasMilitaryDepartment') {
                  const russianSubjectName = Object.keys(subjectNameMapping).find(key => subjectNameMapping[key] === subject) || subject;
                  return (
                     <label key={subject} className="checkbox-label">
                        <input
                           type="checkbox"
                           name={russianSubjectName}
                           checked={isChecked}
                           onChange={handleCheckboxChange}
                        />
                        {russianSubjectName}
                     </label>
                  );
               }
               return null;
            })}
         </div>

         <div className="filter-section">
            <label className="label">
               Государственный вуз
               <input
                  type="checkbox"
                  name="isStateUniversity"
                  checked={filters.isStateUniversity}
                  onChange={handleCheckboxChange}
               />
            </label>

            <label className="label voenka">
               Военная кафедра
               <input
                  type="checkbox"
                  name="hasMilitaryDepartment"
                  checked={filters.hasMilitaryDepartment}
                  onChange={handleCheckboxChange}
               />
            </label>
         </div>

         <div className="filter-buttons">
            <button className="filter-button-admin" onClick={handleFilterClick}>
               Применить фильтры
            </button>

         </div>

         <label className="label">
            Введите количество записей для отчета:
            <input
               type="number"
               value={numberOfRecords}
               onChange={(e) => setNumberOfRecords(e.target.value)}
            />
         </label>
         <button className="filter-button-admin" onClick={handleGetInfoClick}>
            Получить информацию по n записям
         </button>

         {selectedUniversity && (
            <div className="edit-form">
               <h2>
                  Поменять данные: {selectedUniversity.university.name} - {selectedUniversity.name}
                  {selectedUniversity.price === 0 ? ' (Бесплатное обучение)' : ` (Платное обучение: ${selectedUniversity.price} руб)`}
               </h2>

               <label>
                  Цена за обучение:
                  <input
                     type="number"
                     value={selectedUniversity.price}
                     onChange={(e) => handleEditInputChange('price', e.target.value)}
                  />
               </label>
               <label>
                  Количество мест:
                  <input
                     type="number"
                     value={selectedUniversity.places}
                     onChange={(e) => handleEditInputChange('places', e.target.value)}
                  />
               </label>

               <label>
                  Проходной балл:
                  <input
                     type="number"
                     value={selectedUniversity.passingScore}
                     onChange={(e) => handleEditInputChange('passingScore', e.target.value)}
                  />
               </label>

               <label>
                  Присутствие военной кафедры:
                  <input
                     type="checkbox"
                     checked={selectedUniversity.hasMilitaryDepartment}
                     onChange={(e) => handleEditInputChange('hasMilitaryDepartment', e.target.checked)}
                  />
               </label>

               <label>
                  Государственный вуз:
                  <input
                     type="checkbox"
                     checked={selectedUniversity.isStateUniversity}
                     onChange={(e) => handleEditInputChange('isStateUniversity', e.target.checked)}
                  />
               </label>
               <button onClick={handleSaveChanges}>Сохранить изменения</button>


            </div>
         )}

         {responseData && <UniversityListAdmin universities={responseData} onEditClick={handleEditClick} />}

      </div>
   );
};

export default AdminPanel;
