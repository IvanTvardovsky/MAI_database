package controllers

import (
	"backend/internal/db"
	"backend/internal/schemas"
	"github.com/gin-gonic/gin"
	"log"
)

func ChangePlace(c *gin.Context) {
	s1 := `SELECT u.id FROM university u WHERE u.name = $1`
	s6 := `SELECT ur.place FROM university_rating ur WHERE ur.university_id = $1`
	s3 := `UPDATE university_rating SET place = place + 1 WHERE id = $1`
	s5 := `UPDATE university_rating SET place = place - 1 WHERE id = $1`
	s4 := `SELECT ur.id FROM university_rating ur WHERE ur.id BETWEEN $1 AND $2 ORDER BY ur.id`
	s2 := `UPDATE university_rating SET place = $1 WHERE university_id = $2`
	database, _ := db.GetDB()
	up := schemas.UniversityPlace{}
	e := c.BindJSON(&up)
	if e != nil {
		log.Fatalln(e)
	}
	uniId := -1
	oldPlace := -1
	row := database.QueryRow(s1, up.UniversityName)
	row.Scan(&uniId)
	row = database.QueryRow(s6, uniId)
	row.Scan(&oldPlace)
	if oldPlace > up.Place {
		r, _ := database.Query(s4, up.Place, oldPlace-1)
		for r.Next() {
			id := 0
			r.Scan(&id)
			database.Exec(s3, id)
		}
	} else if oldPlace < up.Place {
		r, _ := database.Query(s4, oldPlace+1, up.Place)
		for r.Next() {
			id := 0
			r.Scan(&id)
			database.Exec(s5, id)
		}
	}
	database.Exec(s2, up.Place, uniId)
	c.JSON(200, struct {
		Resp string
	}{"OK"})
	return
}
