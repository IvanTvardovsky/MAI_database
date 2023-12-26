package schemas

type University struct {
	Id          int                 `json:"id"`
	IsState     bool                `json:"is_state_university"`
	HasMilitary bool                `json:"has_military_department"`
	Place       int                 `json:"place"`
	Name        string              `json:"name"`
	Programs    []ProgramUniversity `json:"programs"`
}

type UniversityVanyasha struct {
	Id          int    `json:"id"`
	IsState     bool   `json:"is_state_university"`
	HasMilitary bool   `json:"has_military_department"`
	Place       int    `json:"place"`
	Name        string `json:"name"`
}

type ProgramUniversity struct {
	Id           int       `json:"id"`
	Price        int       `json:"price"`
	Places       int       `json:"places"`
	Subjects     []Subject `json:"subjects"`
	PassingScore int       `json:"passingScore"`
	Program
	P_type int `json:"p_type"`
}

type ProgramUniversityVanyasha struct {
	Id           int       `json:"id"`
	Price        int       `json:"price"`
	Places       int       `json:"places"`
	Subjects     []Subject `json:"subjects"`
	PassingScore int       `json:"passingScore"`
	Program
	P_type     int                `json:"p_type"`
	University UniversityVanyasha `json:"university"`
}

type Subject struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	IsChoosable bool   `json:"is_choosable"`
}

type Program struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type FiltersUni struct {
	HasMilitary  *bool
	Subjects     []string
	PaidPlaces   *int
	BudgetPlaces *int
	PaidCost     *int
	Place        *int
	BudgetScore  *int
	PaidScore    *int
	IsState      *bool
}

type UniversityPlace struct {
	UniversityName string
	Place          int
}
