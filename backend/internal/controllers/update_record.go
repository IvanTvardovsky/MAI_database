package controllers

import (
	"backend/internal/db"
	"backend/internal/schemas"
	"github.com/gin-gonic/gin"
	"log"
)

func UpdateUniversity(c *gin.Context) {
	s1 := `UPDATE university SET name=$1,has_military_department=$2,
    is_state_university=$3 WHERE id=$4`
	s2 := `UPDATE university_rating SET place=$1 WHERE university_id=$2`
	s3 := `UPDATE university_program SET places=$1, passing_score=$2 WHERE id=$3`
	s4 := `UPDATE university_program_paid SET price=$1 WHERE id=$2`
	database, _ := db.GetDB()
	updatedData := schemas.ProgramUniversityVanyasha{}
	err := c.BindJSON(&updatedData)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = database.Exec(s1, updatedData.University.Name, updatedData.University.HasMilitary, updatedData.University.IsState,
		updatedData.University.Id)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = database.Exec(s2, updatedData.University.Place, updatedData.University.Id)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = database.Exec(s3, updatedData.Places, updatedData.PassingScore, updatedData.Id)
	log.Println(updatedData.Places, updatedData.PassingScore, updatedData.Id)
	if err != nil {
		log.Fatalln(err)
	}
	if updatedData.P_type == 1 {
		_, err = database.Exec(s4, updatedData.Price, updatedData.Id)
		if err != nil {
			log.Fatalln(err)
		}
	}
	c.JSON(200, struct{ Status string }{"ok"})
	return
}
