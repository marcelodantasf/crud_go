package main

import (
	"bufio"
	"empresa/functions"
	"empresa/settings"
	"fmt"
	"os"
)

func initializeFiles() {
	files := []string{settings.EmployeeFilename, settings.DepartmentFilename, settings.ProjectFilename, settings.IDTrackerFilename}

	for _, filename := range files {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}

func employeeMenu() {
	var op int
	reader := bufio.NewReader(os.Stdin)

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

		reader.ReadString('\n')

		switch op {
		case 1:
			fmt.Println("Criando funcionário...")
			functions.CreateEmployee()
		case 2:
			employees, err := functions.ReadEmployees()
			if err != nil {
				fmt.Println("Erro ao ler funcionários:", err)
				break
			}
			for _, emp := range employees {
				fmt.Printf("Nome: %s %s %s | CPF: %s | Salário: %.2f | Gênero: %s | Nascimento: %s\n",
					emp.FirstName, emp.MiddleNameInitial, emp.LastName,
					emp.CPF, emp.Salary, emp.Gender, emp.DateOfBirth)
			}
		case 3:
			fmt.Print("Digite o CPF do funcionário a ser alterado: ")
			var cpf string
			fmt.Scan(&cpf)
			functions.UpdateEmployee(cpf)

			//fmt.Println("Alterando funcionário...")
		case 4:
			fmt.Println("Excluindo funcionário...")
			fmt.Print("Digite o CPF do funcionário a ser deletado: ")
			var cpf string
			fmt.Scan(&cpf)
			functions.DeleteEmployeeByCPF(cpf)
		case 5:
			fmt.Println("Buscar funcionário por CPF...")
			fmt.Print("Digite o CPF: ")
			var cpf string
			fmt.Scan(&cpf)

			reader.ReadString('\n')

			employee, err := functions.ReadEmployeeByCPF(cpf)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Nome: %s %s %s\n", employee.FirstName, employee.MiddleNameInitial, employee.LastName)
				fmt.Printf("CPF: %s\n", employee.CPF)
				fmt.Printf("Salário: %.2f\n", employee.Salary)
				fmt.Printf("Gênero: %s\n", employee.Gender)
				fmt.Printf("Data de nascimento: %s\n", employee.DateOfBirth)
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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Projetos ---")
		fmt.Println("1: Criar projetos")
		fmt.Println("2: Listar projetos")
		fmt.Println("3: Alterar projeto")
		fmt.Println("4: Excluir projeto")
		fmt.Println("5: Buscar projeto")
		fmt.Println("0: Voltar")
		fmt.Print("Digite a opção: ")

		// Lê a opção do menu
		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		reader.ReadString('\n')

		switch op {
		case 1:
			fmt.Println("Criando projeto...")
			functions.CreateProject()
		case 2:
			fmt.Println("Listando projetos...")
			projects, err := functions.ReadProjects()
			if err != nil {
				fmt.Println("Erro ao ler projetos:", err)
				break
			}
			for _, project := range projects {
				fmt.Printf("ID: %d | Nome: %s | Local: %s\n", project.ID, project.Name, project.Local)
			}
		case 3:
			fmt.Println("Alterando projeto...")
			fmt.Print("Digite o ID do projeto a ser alterado: ")
			var id int
			fmt.Scan(&id)
			functions.UpdateProject(id)
		case 4:
			fmt.Println("Excluindo projeto...")
			fmt.Print("Digite o ID do projeto a ser deletado: ")
			var id int
			fmt.Scan(&id)
			functions.DeleteProjectByID(id)
		case 5:
			fmt.Println("Buscar projeto por ID...")
			fmt.Print("Digite o ID do projeto: ")
			var id int
			fmt.Scan(&id)

			reader.ReadString('\n')

			project, err := functions.ReadProjectByID(id)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("ID: %d | Nome: %s | Local: %s\n", project.ID, project.Name, project.Local)
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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Departamentos ---")
		fmt.Println("1: Criar departamento")
		fmt.Println("2: Listar departamentos")
		fmt.Println("3: Alterar departamento")
		fmt.Println("4: Excluir departamento")
		fmt.Println("5: Buscar departamento")
		fmt.Println("6: Definir gerente")
		fmt.Println("0: Voltar")
		fmt.Print("Digite a opção: ")

		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		reader.ReadString('\n')

		switch op {
		case 1:
			fmt.Println("Criando departamento...")
			err := functions.CreateDepartment()
			if err != nil {
				fmt.Println("Erro ao criar departamento:", err)
			}
		case 2:
			fmt.Println("Listando departamentos...")
		case 3:
			fmt.Println("Alterando departamento...")
			fmt.Print("Digite o ID do departamento a ser alterado: ")
			var id int
			fmt.Scan(&id)
			functions.UpdateDepartment(id)
		case 4:
			fmt.Println("Excluindo departamento...")
			fmt.Print("Digite o ID do departamento a ser alterado: ")
			var id int
			fmt.Scan(&id)
			functions.DeleteDepartmentByID(id)
		case 5:
			fmt.Println("Buscando departamento...")
			fmt.Print("Digite o ID do departamento: ")
			var id int
			fmt.Scan(&id)

			reader.ReadString('\n')

			dept, err := functions.ReadDepartmentByID(id)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("ID: %d | Nome: %s\n", dept.ID, dept.Name)
			}
		case 6:
			fmt.Print("Digite o ID do departamento para definir o gerente: ")
			var id int
			fmt.Scan(&id)
			functions.SelectManager(id)
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
