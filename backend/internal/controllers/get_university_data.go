package controllers

import (
	"backend/internal/db"
	"backend/internal/schemas"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetUniversityData(c *gin.Context) {
	const s1 = `SELECT u.id, u.is_state_university, u.has_military_department, u.name, ur.place FROM university u, university_rating ur  WHERE ur.id = $1 AND u.id = $1`
	const s2 = `SELECT up.id, up.passing_score, up.places, p.name, p.code  FROM university_program up JOIN program p ON up.program_id = p.id WHERE up.university_id = $1`
	const s3 = `SELECT upb.id  FROM university_program up JOIN university_program_budget upb ON up.id = upb.id WHERE up.id = $1`
	const s4 = `SELECT upp.price FROM university_program up JOIN university_program_paid upp ON up.id = upp.id WHERE up.id = $1`
	const s5 = `SELECT ups.is_choosable, s.name FROM university_program_subjects ups JOIN subjects s ON ups.subject_id = s.id WHERE ups.program_id = $1`
	uni := schemas.University{}
	database, _ := db.GetDB()
	uni_id, _ := strconv.Atoi(c.Param("id"))
	row := database.QueryRow(s1, uni_id)
	row.Scan(&uni.Id, &uni.IsState, &uni.HasMilitary, &uni.Name, &uni.Place)
	rows, _ := database.Query(s2, uni_id)
	for rows.Next() {
		p := schemas.ProgramUniversity{}
		rows.Scan(&p.Id, &p.PassingScore, &p.Places, &p.Name, &p.Code)
		row = database.QueryRow(s3, p.Id)
		log.Println(p.Id)
		t := 0
		err := row.Scan(&t)
		if err != nil {
			row = database.QueryRow(s4, p.Id)
			row.Scan(&p.Price)
			p.P_type = 1
		} else {
			p.P_type = 0
		}
		rows, _ := database.Query(s5, p.Id)
		for rows.Next() {
			s := schemas.Subject{}
			rows.Scan(&s.IsChoosable, &s.Name)
			p.Subjects = append(p.Subjects, s)
		}
		uni.Programs = append(uni.Programs, p)
		rows.Close()

	}
	rows.Close()
	c.JSON(200, uni)
	return
}
