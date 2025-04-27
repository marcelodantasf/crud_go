package main

/*import(
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
) */

type Employee struct {
	firstName         string
	middleNameInitial string
	lastName          string
	cpf               string
	address           string
	salary            float64
	gender            string
	dateOfBirth       string
	department        *Department
	projects          []*Project
}

type Project struct {
	name  string
	id    int
	local string
}

type Department struct {
	id       int
	name     string
	manager  *Employee
	projects []*Project
}

const employee_filename = "employee.csv"
const project_filename = "project.csv"
const department_filename = "department.csv"

func createEmployee(Employee Employee) //falta add a exception
func readEmployees() ([]Employee, error)
func createProject(Project Project) //+exception
func readProjects() ([]Project, error)
func readProjectByID(id int) (*Project, error)
