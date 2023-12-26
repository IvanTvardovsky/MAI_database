package controllers

import (
	"backend/internal/db"
	"backend/internal/schemas"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"io"
	"net/http"
	"os"
	"strconv"
)

func GetRatingJSON(c *gin.Context) {
	database, _ := db.GetDB()
	s := `SELECT u.name, ur.place FROM university u JOIN university_rating ur ON ur.university_id = u.id ORDER BY place ASC`
	rows, _ := database.Query(s)
	defer rows.Close()
	places := make([]schemas.UniversityPlace, 0)
	for rows.Next() {
		t := schemas.UniversityPlace{}
		rows.Scan(&t.UniversityName, &t.Place)
		places = append(places, t)
	}
	c.IndentedJSON(http.StatusOK, places)
}

func GetRatingREPORT(c *gin.Context) {
	database, _ := db.GetDB()
	s := `SELECT u.name, ur.place FROM university u JOIN university_rating ur ON ur.university_id = u.id ORDER BY place ASC`
	rows, _ := database.Query(s)
	defer rows.Close()
	places := make([]schemas.UniversityPlace, 0)
	for rows.Next() {
		t := schemas.UniversityPlace{}
		rows.Scan(&t.UniversityName, &t.Place)
		places = append(places, t)
	}
	createPDFRating(places)
	c.Writer.Header().Set("Content-type", "application/pdf")
	f, err := os.Open("./report_rating.pdf")
	if err != nil {
		fmt.Println(err)
	}
	if _, err := io.Copy(c.Writer, f); err != nil {
		fmt.Println(err)
	}
	f.Close()
	os.Remove("./report_rating.pdf")
	return
}

func createPDFRating(data []schemas.UniversityPlace) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(true)
	pdf.SetTopMargin(10)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.SetFont("Helvetica", "", 12)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	for _, e := range data {
		pdf.Write(7, tr("Название университета: "+e.UniversityName+"\n"))
		place := strconv.Itoa(e.Place)
		pdf.Write(7, tr("Место: "+place+".\n\n"))
	}

	_ = pdf.OutputFileAndClose("report_rating.pdf")
}
