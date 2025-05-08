package models

type Employee struct {
	FirstName         string
	MiddleNameInitial string
	LastName          string
	CPF               string
	Address           string
	Salary            float64
	Gender            string
	DateOfBirth       string
	Department        *Department
	Projects          []*Project
}

type Department struct {
	ID       int
	Name     string
	Manager  *Employee
	Projects []*Project
}

type Project struct {
	ID    int
	Name  string
	Local string
}
