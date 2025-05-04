package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func initializeFiles() {
	files := []string{employeeFilename, departmentFilename, projectFilename}

	for _, filename := range files {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}

func createEmployee(employee Employee) error {
	file, err := os.OpenFile(employeeFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s,%s,%s,%s,%f,%s,%s\n", employee.firstName, employee.middleNameInitial, employee.lastName, employee.cpf, employee.salary, employee.gender, employee.dateOfBirth))
	return err
}

// falta add a exception

func readEmployees() ([]Employee, error) {
	file, err := os.Open(employeeFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var employees []Employee
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) < 7 {
			continue // Linha mal formatada
		}

		salary, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			continue
		}

		employee := Employee{
			firstName:         fields[0],
			middleNameInitial: fields[1],
			lastName:          fields[2],
			cpf:               fields[3],
			salary:            salary,
			gender:            fields[5],
			dateOfBirth:       fields[6],
		}
		employees = append(employees, employee)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func readEmployeeByCPF(cpf string) (*Employee, error) {
	employees, err := readEmployees()
	if err != nil {
		return nil, err
	}

	for _, emp := range employees {
		if emp.cpf == cpf {
			return &emp, nil
		}
	}

	return nil, fmt.Errorf("Funcionário com CPF %s não encontrado", cpf)
}

func readProjects() ([]Project, error) {
	file, err := os.Open(projectFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var projects []Project
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			continue
		}

		id, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}

		project := Project{
			id:    id,
			name:  fields[1],
			local: fields[2],
		}
		projects = append(projects, project)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func readProjectByID(id int) (*Project, error) {
	projects, err := readProjects()
	if err != nil {
		return nil, err
	}

	for _, proj := range projects {
		if proj.id == id {
			return &proj, nil
		}
	}

	return nil, fmt.Errorf("Projeto com ID %d não encontrado", id)
}

func readDepartments() ([]Department, error) {
	file, err := os.Open(departmentFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var departments []Department
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			continue
		}

		id, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}

		// Busca o gerente pelo CPF
		manager, _ := readEmployeeByCPF(fields[2]) // se não encontrar, será nil

		dept := Department{
			id:      id,
			name:    fields[1],
			manager: manager,
		}
		departments = append(departments, dept)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func readDepartmentByID(id int) (*Department, error) {
	depts, err := readDepartments()
	if err != nil {
		return nil, err
	}

	for _, dept := range depts {
		if dept.id == id {
			return &dept, nil
		}
	}

	return nil, fmt.Errorf("Departamento com ID %d não encontrado", id)
}

/*func createProject(Project Project) //+exception
func readProjects() ([]Project, error)
func readProjectByID(id int) (*Project, error)*/

func employeeMenu() {
	var op int
	for {
		fmt.Println("\n--- Funcionários ---")
		fmt.Println("1: Criar funcionário")
		fmt.Println("2: Listar funcionários")
		fmt.Println("3: Alterar funcionário")
		fmt.Println("4: Excluir funcionário")
		fmt.Println("5: Buscar funcionário")
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
			var emp1 Employee
			emp1.firstName = "kelvin"
			emp1.middleNameInitial = "L"
			emp1.lastName = "Rodrigues"
			emp1.address = "av 13 de maio"
			emp1.cpf = "61824604319"
			emp1.salary = 38813.32
			emp1.gender = "M"
			createEmployee(emp1)

		case 2:
			employees, err := readEmployees()
			if err != nil {
				fmt.Println("Erro ao ler funcionários:", err)
				break
			}
			for _, emp := range employees {
				fmt.Printf("Nome: %s %s %s | CPF: %s | Salário: %.2f | Gênero: %s | Nascimento: %s\n",
					emp.firstName, emp.middleNameInitial, emp.lastName,
					emp.cpf, emp.salary, emp.gender, emp.dateOfBirth)
			}
		case 3:
			fmt.Println("Alterando funcionário...")
		case 4:
			fmt.Println("Excluindo funcionário...")
		case 5:
			fmt.Println("Buscar funcionário por CPF...")
			fmt.Print("Digite o CPF: ")
			var cpf string
			fmt.Scan(&cpf)

			employee, err := readEmployeeByCPF(cpf)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Nome: %s %s %s\n", employee.firstName, employee.middleNameInitial, employee.lastName)
				fmt.Printf("CPF: %s\n", employee.cpf)
				fmt.Printf("Salário: %.2f\n", employee.salary)
				fmt.Printf("Gênero: %s\n", employee.gender)
				fmt.Printf("Data de nascimento: %s\n", employee.dateOfBirth)
			}
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
		fmt.Println("5: Buscar projeto")
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
			projects, err := readProjects()
			if err != nil {
				fmt.Println("Erro ao ler projetos:", err)
				break
			}
			for _, project := range projects {
				fmt.Printf("ID: %d | Nome: %s | Local: %s\n", project.id, project.name, project.local)
			}
		case 3:
			fmt.Println("Alterando projeto...")
		case 4:
			fmt.Println("Excluindo projeto...")
		case 5:
			fmt.Println("Buscar projeto por ID...")
			fmt.Print("Digite o ID do projeto: ")
			var id int
			fmt.Scan(&id)

			project, err := readProjectByID(id)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("ID: %d | Nome: %s | Local: %s\n", project.id, project.name, project.local)
			}
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
		fmt.Println("5: Buscar departamento")
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
		case 5:
			fmt.Println("Buscando departamento...")
			fmt.Print("Digite o ID do departamento: ")
			var id int
			fmt.Scan(&id)

			dept, err := readDepartmentByID(id)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("ID: %d | Nome: %s\n", dept.id, dept.name)
			}
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
