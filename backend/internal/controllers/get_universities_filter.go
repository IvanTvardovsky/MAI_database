package controllers

import (
	"backend/internal/db"
	"backend/internal/schemas"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//func buildRequest(c *gin.Context) {
//
//}

func getQueryParams(c *gin.Context, p *schemas.FiltersUni) {
	Military, has := c.GetQuery("military")
	if has {
		MilitaryParsed, _ := strconv.ParseBool(Military)
		p.HasMilitary = &MilitaryParsed
	}
	PaidCost, has := c.GetQuery("paid_cost")
	if has {
		t, _ := strconv.Atoi(PaidCost)
		p.PaidCost = &t
	}
	PaidPlaces, has := c.GetQuery("paid_places")
	if has {
		t, _ := strconv.Atoi(PaidPlaces)
		p.PaidPlaces = &t
	}
	BudgetPlaces, has := c.GetQuery("budget_places")
	if has {
		t, _ := strconv.Atoi(BudgetPlaces)
		p.BudgetPlaces = &t
	}
	subjects, has := c.GetQuery("subjects")
	if has {
		t := strings.Split(subjects, ",")
		p.Subjects = t
	}
	place, has := c.GetQuery("place")
	if has {
		t, _ := strconv.Atoi(place)
		p.Place = &t
	}
	paid_score, has := c.GetQuery("paid_score")
	if has {
		t, _ := strconv.Atoi(paid_score)
		p.PaidScore = &t
	}
	budget_score, has := c.GetQuery("budget_score")
	if has {
		t, _ := strconv.Atoi(budget_score)
		p.BudgetScore = &t
	}
	State, has := c.GetQuery("state")
	if has {
		StateParsed, _ := strconv.ParseBool(State)
		p.IsState = &StateParsed
	}
}

func postHistory(f schemas.FiltersUni) {
	database, _ := db.GetDB()
	s1 := `INSERT INTO history(date, budget_places, paid_places,budget_passing_score, paid_passing_score, paid_cost, is_state, has_military, subjects, place) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	sub := strings.Join(f.Subjects, ",")
	var subP *string = nil
	if len(sub) != 0 {
		subP = &sub
	}
	if f.BudgetPlaces == nil && f.PaidPlaces == nil && f.BudgetScore == nil && f.PaidScore == nil && f.PaidScore == nil && f.IsState == nil && f.HasMilitary == nil && len(sub) == 0 && f.Place == nil {
		return
	}
	_, err := database.Exec(s1, time.Now(), f.BudgetPlaces, f.PaidPlaces, f.BudgetScore, f.PaidScore, f.PaidScore, f.IsState, f.HasMilitary, subP, f.Place)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetUniversitiesDataFilter(c *gin.Context) {
	database, _ := db.GetDB()
	f := schemas.FiltersUni{}
	getQueryParams(c, &f)
	postHistory(f)
	const s10 = `SELECT u.id FROM university u`
	ret := []schemas.University{}
	rows, _ := database.Query(s10)
	defer rows.Close()
	for rows.Next() {
		const s1 = `SELECT u.id, u.is_state_university, u.has_military_department, u.name, ur.place FROM university u, university_rating ur  WHERE ur.university_id = $1 AND u.id = $1`
		const s2 = `SELECT up.id, up.passing_score, up.places, p.name, p.code  FROM university_program up JOIN program p ON up.program_id = p.id WHERE up.university_id = $1`
		const s3 = `SELECT upb.id  FROM university_program up JOIN university_program_budget upb ON up.id = upb.id WHERE up.id = $1`
		const s4 = `SELECT upp.price FROM university_program up JOIN university_program_paid upp ON up.id = upp.id WHERE up.id = $1`
		const s5 = `SELECT ups.is_choosable, s.name FROM university_program_subjects ups JOIN subjects s ON ups.subject_id = s.id WHERE ups.program_id = $1`
		uni := schemas.University{}
		uni_id := 0
		rows.Scan(&uni_id)
		row := database.QueryRow(s1, uni_id)
		row.Scan(&uni.Id, &uni.IsState, &uni.HasMilitary, &uni.Name, &uni.Place)
		if f.Place != nil && uni.Place > *f.Place {
			continue
		}
		rows, _ := database.Query(s2, uni_id)

	p:
		for rows.Next() {
			p := schemas.ProgramUniversity{}
			rows.Scan(&p.Id, &p.PassingScore, &p.Places, &p.Name, &p.Code)
			row = database.QueryRow(s3, p.Id)
			t := 0
			err := row.Scan(&t)
			if err != nil {
				row = database.QueryRow(s4, p.Id)
				row.Scan(&p.Price)
				p.P_type = 1
			} else {
				p.P_type = 0
			}
			if f.BudgetPlaces != nil && p.P_type == 0 && p.Places < *f.BudgetPlaces {
				continue
			}
			if f.PaidPlaces != nil && p.P_type == 1 && p.Places < *f.PaidPlaces {
				continue
			}
			if f.PaidCost != nil && p.P_type == 1 && p.Price < *f.PaidCost {
				continue
			}
			if f.PaidScore != nil && p.P_type == 1 && p.PassingScore > *f.PaidScore {
				continue
			}
			if f.BudgetScore != nil && p.P_type == 0 && p.PassingScore > *f.BudgetScore {
				continue
			}
			rows, _ := database.Query(s5, p.Id)
			subjectsProgram := make([]schemas.Subject, 0)
			for rows.Next() {
				s := schemas.Subject{}
				rows.Scan(&s.IsChoosable, &s.Name)
				subjectsProgram = append(subjectsProgram, s)
			}
			if f.Subjects != nil {
				for _, sub := range f.Subjects {
					flag := false
					for _, sp := range subjectsProgram {
						//log.Println(el)
						if sub == sp.Name {
							flag = true
						}
					}
					if !flag {
						continue p
						rows.Close()
					}
				}
			}
			p.Subjects = subjectsProgram
			uni.Programs = append(uni.Programs, p)
			rows.Close()

		}
		rows.Close()
		if len(uni.Programs) != 0 {
			ret = append(ret, uni)
		}

	}
	rows.Close()
	newStyle := []schemas.ProgramUniversityVanyasha{}
	for _, uni := range ret {
		for _, pr := range uni.Programs {
			t := schemas.ProgramUniversityVanyasha{}
			t.University.Id = uni.Id
			t.University.Name = uni.Name
			t.University.HasMilitary = uni.HasMilitary
			t.University.IsState = uni.IsState
			t.University.Place = uni.Place
			t.Id = pr.Id
			t.Price = pr.Price
			t.Places = pr.Places
			t.Subjects = pr.Subjects
			t.Name = pr.Name
			t.Code = pr.Code
			t.P_type = pr.P_type
			t.PassingScore = pr.PassingScore
			newStyle = append(newStyle, t)
		}
	}
	c.IndentedJSON(http.StatusOK, newStyle)

}
