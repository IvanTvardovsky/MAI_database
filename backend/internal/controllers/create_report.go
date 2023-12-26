package controllers

import (
	"backend/internal/db"
	"backend/internal/schemas"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetReport(c *gin.Context) {
	counter := 0
	database, _ := db.GetDB()
	v, has := c.GetQuery("n")
	value, _ := strconv.Atoi(v)
	report := schemas.Report{}
	s1 := `SELECT date, budget_places, paid_places,budget_passing_score, paid_passing_score, paid_cost, is_state, has_military, subjects FROM history ORDER BY date`
	if has {
		s1 += ` LIMIT $1`
	}
	var rows *sql.Rows
	var err error
	if has {
		rows, err = database.Query(s1, value)
	} else {
		rows, err = database.Query(s1)
	}
	if err != nil {
		log.Fatalln(err)
	}
	dates := make([]time.Time, value)
	for rows.Next() {
		f := schemas.FiltersUni{}
		sub := ""
		rows.Scan(&dates[counter], &f.BudgetPlaces, &f.PaidPlaces, &f.BudgetScore, &f.PaidScore, &f.PaidScore, &f.IsState, &f.HasMilitary, &sub)
		t := make([]string, 0)
		if sub != "" {
			t = strings.Split(sub, ",")
		}

		f.Subjects = t
		report.Records = append(report.Records, f)
		counter++
		if counter == value {
			break
		}
	}
	c.Writer.Header().Set("Content-type", "application/pdf")
	createPDF(report, dates)
	f, err := os.Open("./report.pdf")
	if err != nil {
		fmt.Println(err)
	}
	if _, err := io.Copy(c.Writer, f); err != nil {
		fmt.Println(err)
	}
	//c.File("./report.pdf")
	f.Close()
	err = os.Remove("./report.pdf")
	return
}

func createPDF(r schemas.Report, d []time.Time) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(true)
	pdf.SetTopMargin(10)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.SetFont("Helvetica", "", 12)
	t := strconv.Itoa(len(r.Records))
	count := len(r.Records)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	pdf.Write(7, tr("Всего запрошено записей: "+t+".\n\n"))
	sumBudget := 0
	sumPaid := 0
	sumPlacesPaid := 0
	sumPlacesBudget := 0

	for i, e := range r.Records {
		sb := strings.Builder{}
		sb.WriteString("Дата запроса: ")
		//tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
		sb.WriteString(d[i].Format("2006-01-02") + ". ")
		sb.WriteString("Используемые параметры: ")
		if e.BudgetPlaces != nil {
			sumPlacesBudget += *e.BudgetPlaces
			s := strconv.Itoa(*e.BudgetPlaces)
			sb.WriteString("Количество бюджетных мест = " + s + ". ")
		}
		if e.PaidPlaces != nil {
			sumPlacesPaid += *e.PaidPlaces
			s := strconv.Itoa(*e.PaidPlaces)
			sb.WriteString("Количество платных мест = " + s + ". ")
		}
		if e.PaidScore != nil {
			sumPaid += *e.PaidScore
			s := strconv.Itoa(*e.PaidScore)
			sb.WriteString("Проходной балл на платку = " + s + ". ")
		}
		if e.BudgetScore != nil {
			sumBudget += *e.BudgetScore
			s := strconv.Itoa(*e.BudgetScore)
			sb.WriteString("Проходной балл на бюджет = " + s + ". ")
		}
		if e.IsState != nil {
			s := ""
			if *e.IsState {
				s = "Да"
			} else {
				s = "Нет"
			}
			sb.WriteString("Должен ли вуз быть государственным: " + s + ". ")
		}
		if e.HasMilitary != nil {
			s := ""
			if *e.HasMilitary {
				s = "Да"
			} else {
				s = "Нет"
			}
			sb.WriteString("Должен ли вуз иметь военную кафедру: " + s + ". ")
		}
		if e.Place != nil {
			s := strconv.Itoa(*e.Place)
			sb.WriteString("Место университета (не ниже данного): " + s + ". ")
		}
		if len(e.Subjects) != 0 {
			s := strings.Join(e.Subjects, ",")
			sb.WriteString("Предметы для сдачи: " + s + ".")
		}
		sb.WriteString("\n\n")
		//pdf.CellFormat(0, 10, tr(sb.String()), "", 1, "", false, 0, "")
		pdf.Write(7, tr(sb.String()))
	}
	srSumBudget := fmt.Sprintf("%f", math.Round(float64(sumBudget)/float64(count)))
	srSumPaid := fmt.Sprintf("%f", math.Round(float64(sumPaid)/float64(count)))
	srSumBudgetPlaces := fmt.Sprintf("%f", math.Round(float64(sumPlacesBudget)/float64(count)))
	srSumPaidPlaces := fmt.Sprintf("%f", math.Round(float64(sumPlacesPaid)/float64(count)))
	pdf.Write(7, tr("Среднее запрашиваемое количество бюджетных мест: "+srSumBudgetPlaces+".\n"))
	pdf.Write(7, tr("Среднее запрашиваемое количество платных мест: "+srSumPaidPlaces+".\n"))
	pdf.Write(7, tr("Среднее запрашиваемое количество бюджетных баллов: "+srSumBudget+".\n"))
	pdf.Write(7, tr("Среднее запрашиваемое количество платных баллов: "+srSumPaid+".\n"))
	err := pdf.OutputFileAndClose("report.pdf")
	if err != nil {
		log.Fatalln(err)
	}
}
