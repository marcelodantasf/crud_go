package main

import (
	"fmt"
	"os"
	//"strconv"
)

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

const (
	employeeFilename   = "employee.txt"
	projectFilename    = "project.txt"
	departmentFilename = "department.txt"
)

func createEmployee(Employee Employee) {
}

// falta add a exception
/*func readEmployees() ([]Employee, error)
func createProject(Project Project) //+exception
func readProjects() ([]Project, error)
func readProjectByID(id int) (*Project, error)*/

func initializeFiles() {
	employee_file, err := os.Create(employeeFilename)
	if err != nil {
		panic(0)
	}
	defer employee_file.Close()

	department_file, err := os.Create(departmentFilename)
	if err != nil {
		panic(0)
	}
	defer department_file.Close()

	project_file, err := os.Create(projectFilename)
	if err != nil {
		panic(0)
	}
	defer project_file.Close()
}

func employeeMenu() {
	var op int
	for {
		fmt.Println("\n--- Funcionários ---")
		fmt.Println("1: Criar funcionário")
		fmt.Println("2: Listar funcionários")
		fmt.Println("3: Alterar funcionário")
		fmt.Println("4: Excluir funcionário")
		fmt.Println("0: Voltar")
		fmt.Print("Digite a opção: ")

		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Println("Opção inválida. Tente novamente.")
			continue
		}

		switch op {
		case 1:
			// TODO: implementar
			fmt.Println("Criando funcionário...")
		case 2:
			fmt.Println("Listando funcionários...")
		case 3:
			fmt.Println("Alterando funcionário...")
		case 4:
			fmt.Println("Excluindo funcionário...")
		case 0:
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func projectMenu() {
	var op int
	for {
		fmt.Println("\n--- Projetos ---")
		fmt.Println("1: Criar projetos")
		fmt.Println("2: Listar projetos")
		fmt.Println("3: Alterar projeto")
		fmt.Println("4: Excluir projeto")
		fmt.Println("0: Voltar")
		fmt.Print("Digite a opção: ")

		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch op {
		case 1:
			// TODO: implementar
			fmt.Println("Criando projeto...")
		case 2:
			fmt.Println("Listando projetos...")
		case 3:
			fmt.Println("Alterando projeto...")
		case 4:
			fmt.Println("Excluindo projeto...")
		case 0:
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func departmentMenu() {
	var op int
	for {
		fmt.Println("\n--- Departamentos ---")
		fmt.Println("1: Criar departamento")
		fmt.Println("2: Listar departamento")
		fmt.Println("3: Alterar departamento")
		fmt.Println("4: Excluir departamento")
		fmt.Println("0: Voltar")
		fmt.Print("Digite a opção: ")

		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch op {
		case 1:
			fmt.Println("Criando departamento...")
		case 2:
			fmt.Println("Listando departamentos...")
		case 3:
			fmt.Println("Alterando departamento...")
		case 4:
			fmt.Println("Excluindo departamento...")
		case 0:
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func mainMenu() {
	var option int

	for {
		fmt.Println("\n--- Selecione uma das entidades para operar ---")
		fmt.Println("1: Funcionário")
		fmt.Println("2: Projetos")
		fmt.Println("3: Departamentos")
		fmt.Println("0: Sair")
		fmt.Print("Digite a opção: ")

		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		switch option {
		case 1:
			employeeMenu()
		case 2:
			projectMenu()
		case 3:
			departmentMenu()
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func main() {
	initializeFiles()
	mainMenu()
}
