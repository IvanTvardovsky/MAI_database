import React, {useEffect, useState} from 'react';
import "./../css/App.css"
import queryString from 'query-string';
import UniversityList from './UniversityList';

const FilterPanel = () => {
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

   const handleInputChange = (e) => {
      const {name, value} = e.target;
      setFilters((prevFilters) => ({
         ...prevFilters,
         [name]: value,
      }));
   };

   const handleCheckboxChange = (e) => {
      const {name, checked} = e.target;
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
   }, [queryStringResult]);

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
            <button className="filter-button" onClick={handleFilterClick}>
               Применить фильтры
            </button>

         </div>

         {responseData && <UniversityList universities={responseData}/>}
      </div>
   );
};

export default FilterPanel;