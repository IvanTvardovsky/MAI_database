package user

type UniversityProgram struct {
	ID                    int
	PassingScore          int
	BudgetPlaces          int
	PaidPlaces            int
	RatingPlace           int
	Price                 int
	Subjects              []int
	HasMilitaryDepartment bool
	IsStateUniversity     bool
	University            string
}
