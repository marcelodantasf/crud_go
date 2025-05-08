package functions

import (
	"bufio"
	"empresa/common"
	"empresa/io"
	"empresa/models"
	"empresa/settings"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateEmployee() error {
	var emp models.Employee
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o primeiro nome: ")
	emp.FirstName, _ = reader.ReadString('\n')
	emp.FirstName = strings.TrimSpace(emp.FirstName)

	fmt.Print("Digite a inicial do nome do meio: ")
	emp.MiddleNameInitial, _ = reader.ReadString('\n')
	emp.MiddleNameInitial = strings.TrimSpace(emp.MiddleNameInitial)

	fmt.Print("Digite o último nome: ")
	emp.LastName, _ = reader.ReadString('\n')
	emp.LastName = strings.TrimSpace(emp.LastName)

	fmt.Print("Digite o CPF: ")
	emp.CPF, _ = reader.ReadString('\n')
	emp.CPF = strings.TrimSpace(emp.CPF)

	fmt.Print("Digite o endereço: ")
	emp.Address, _ = reader.ReadString('\n')
	emp.Address = strings.TrimSpace(emp.Address)

	fmt.Print("Digite o salário: ")
	salaryStr, _ := reader.ReadString('\n')
	salaryStr = strings.TrimSpace(salaryStr)
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		return fmt.Errorf("salário inválido")
	}
	emp.Salary = salary

	fmt.Print("Digite o gênero (M/F): ")
	emp.Gender, _ = reader.ReadString('\n')
	emp.Gender = strings.TrimSpace(emp.Gender)

	fmt.Print("Data de nascimento (DD/MM/AAAA): ")
	emp.DateOfBirth, _ = reader.ReadString('\n')
	emp.DateOfBirth = strings.TrimSpace(emp.DateOfBirth)

	fmt.Print("Digite o ID do departamento: ")
	deptIDStr, _ := reader.ReadString('\n')
	deptIDStr = strings.TrimSpace(deptIDStr)
	deptID, err := strconv.Atoi(deptIDStr)
	if err != nil {
		return fmt.Errorf("ID do departamento inválido")
	}
	dept, err := ReadDepartmentByID(deptID)
	if err != nil {
		return err
	}
	emp.Department = dept

	fmt.Print("Digite o(s) ID(s) do(s) projeto(s) separado(s) por vírgula: ")
	projIDsStr, _ := reader.ReadString('\n')
	projIDsStr = strings.TrimSpace(projIDsStr)
	projIDs := strings.Split(projIDsStr, ",")
	for _, idStr := range projIDs {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		projID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("ID de projeto inválido:", idStr)
			continue
		}
		proj, err := ReadProjectByID(projID)
		if err != nil {
			fmt.Println("Projeto não encontrado:", projID)
			continue
		}
		emp.Projects = append(emp.Projects, proj)
	}

	return io.SaveEmployeeToFile(emp)
}

func UpdateEmployee(cpf string) error {
	employees, err := ReadEmployees()
	if err != nil {
		return fmt.Errorf("erro ao ler funcionários: %v", err)
	}

	var emp *models.Employee
	var index int
	for i, e := range employees {
		if e.CPF == cpf {
			emp = &e
			index = i
			break
		}
	}

	if emp == nil {
		return fmt.Errorf("funcionário com CPF %s não encontrado", cpf)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		var option int

		fmt.Printf("\nEditando funcionário: %s, %s, %s\n", emp.FirstName, emp.MiddleNameInitial, emp.LastName)
		fmt.Println("--- Selecione o campo a ser editado ---")
		fmt.Println("1: Nome")
		fmt.Println("2: Inicial do Nome do Meio")
		fmt.Println("3: Sobrenome")
		fmt.Println("4: Endereço")
		fmt.Println("5: Salário")
		fmt.Println("6: Gênero")
		fmt.Println("7: Data de Nascimento")

		fmt.Println("0: Salvar e sair")
		fmt.Print("Digite a opção: ")

		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		reader.ReadString('\n')

		switch option {
		case 1:
			fmt.Print("Novo primeiro nome: ")
			text, _ := reader.ReadString('\n')
			emp.FirstName = strings.TrimSpace(text)
		case 2:
			fmt.Print("Nova inicial do meio: ")
			text, _ := reader.ReadString('\n')
			emp.MiddleNameInitial = strings.TrimSpace(text)
		case 3:
			fmt.Print("Novo sobrenome: ")
			text, _ := reader.ReadString('\n')
			emp.LastName = strings.TrimSpace(text)
		case 4:
			fmt.Print("Novo endereço: ")
			text, _ := reader.ReadString('\n')
			emp.Address = strings.TrimSpace(text)
		case 5:
			fmt.Print("Novo salário: ")
			fmt.Scan(&emp.Salary)
		case 6:
			fmt.Print("Novo gênero: ")
			text, _ := reader.ReadString('\n')
			emp.Gender = strings.TrimSpace(text)
		case 7:
			fmt.Print("Nova data de nascimento: ")
			text, _ := reader.ReadString('\n')
			emp.DateOfBirth = strings.TrimSpace(text)
		case 8:
			fmt.Print("Digite o novo ID do departamento: ")
			deptIDStr, _ := reader.ReadString('\n')
			deptIDStr = strings.TrimSpace(deptIDStr)
			deptID, err := strconv.Atoi(deptIDStr)
			if err != nil {
				return fmt.Errorf("ID do departamento inválido")
			}
			dept, err := ReadDepartmentByID(deptID)
			if err != nil {
				return err
			}
			emp.Department = dept
		case 9:
			fmt.Print("Digite o(s) novo(s) ID(s) dos projeto(s) separado(s) por vírgula: ")
			projIDsStr, _ := reader.ReadString('\n')
			projIDsStr = strings.TrimSpace(projIDsStr)
			projIDs := strings.Split(projIDsStr, ",")
			for _, idStr := range projIDs {
				idStr = strings.TrimSpace(idStr)
				if idStr == "" {
					continue
				}
				projID, err := strconv.Atoi(idStr)
				if err != nil {
					fmt.Println("ID de projeto inválido:", idStr)
					continue
				}
				proj, err := ReadProjectByID(projID)
				if err != nil {
					fmt.Println("Projeto não encontrado:", projID)
					continue
				}
				emp.Projects = append(emp.Projects, proj)
			}
		case 0:
			employees[index] = *emp
			err := io.SaveEmployeesToFile(employees)
			if err != nil {
				return fmt.Errorf("erro ao salvar funcionário: %v", err)
			}
			fmt.Println("Funcionário atualizado com sucesso.")
			return nil
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func DeleteEmployeeByCPF(cpf string) error {
	employees, err := ReadEmployees()
	if err != nil {
		return fmt.Errorf("Erro ao ler funcionários: %v", err)
	}

	found := false
	var updated []models.Employee
	for _, emp := range employees {
		if emp.CPF == cpf {
			found = true
			continue
		}
		updated = append(updated, emp)
	}

	if !found {
		return fmt.Errorf("Funcionário com CPF %s não encontrado", cpf)
	}

	return io.SaveEmployeesToFile(updated)
}

func ReadEmployeeByCPF(cpf string) (*models.Employee, error) {
	employees, err := ReadEmployees()
	if err != nil {
		return nil, err
	}

	for i := range employees {
		if employees[i].CPF == cpf {
			return &employees[i], nil
		}
	}

	return nil, fmt.Errorf("Funcionário com CPF %s não encontrado", cpf)
}

func ReadEmployees() ([]models.Employee, error) {
	file, err := os.Open(settings.EmployeeFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var employees []models.Employee
	scanner := bufio.NewScanner(file)

	var emp models.Employee
	lineCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if lineCount >= 7 {
				employees = append(employees, emp)
				emp = models.Employee{}
				lineCount = 0
			}
			continue
		}

		switch {
		case strings.HasPrefix(line, "CPF:"):
			emp.CPF = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "CPF:"), ","))
			lineCount++
		case strings.HasPrefix(line, "Primeiro nome:"):
			emp.FirstName = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "Primeiro nome:"), ","))
			lineCount++
		case strings.HasPrefix(line, "Inicial do nome do meio:"):
			emp.MiddleNameInitial = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "Inicial do nome do meio:"), ","))
			lineCount++
		case strings.HasPrefix(line, "Último nome:"):
			emp.LastName = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "Último nome:"), ","))
			lineCount++
		case strings.HasPrefix(line, "Endereço:"):
			emp.Address = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "Endereço:"), ","))
			lineCount++
		case strings.HasPrefix(line, "Salário:"):
			salStr := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "Salário:"), ","))
			sal, err := strconv.ParseFloat(salStr, 64)
			if err == nil {
				emp.Salary = sal
				lineCount++
			}
		case strings.HasPrefix(line, "Gênero:"):
			emp.Gender = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "Gênero:"), ","))
			lineCount++
		case strings.HasPrefix(line, "Data de nascimento:"):
			emp.DateOfBirth = strings.TrimSpace(strings.TrimPrefix(line, "Data de nascimento:"))
			lineCount++
		case strings.HasPrefix(line, "ID do departamento:"):
			deptIDStr := strings.TrimSpace(strings.TrimPrefix(line, "ID do departamento:"))
			if deptIDStr != "" {
				deptID, err := strconv.Atoi(deptIDStr)
				if err == nil {
					dept, err := ReadDepartmentByID(deptID)
					if err == nil {
						emp.Department = dept
					}
				}
			}
		case strings.HasPrefix(line, "ID(s) do(s) projeto(s):"):
			idStr := strings.TrimSpace(strings.TrimPrefix(line, "ID(s) do(s) projeto(s):"))
			ids := strings.Split(idStr, ";")
			for _, s := range ids {
				s = strings.TrimSpace(s)
				if s == "" {
					continue
				}
				pid, err := strconv.Atoi(s)
				if err == nil {
					proj, err := ReadProjectByID(pid)
					if err == nil {
						emp.Projects = append(emp.Projects, proj)
					}
				}
			}
		}
	}

	if lineCount >= 7 {
		employees = append(employees, emp)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func CreateDepartment() error {
	var dep models.Department
	reader := bufio.NewReader(os.Stdin)

	dep.ID, _ = GetNextDepartmentID()

	fmt.Print("Digite o nome do departamento: ")
	dep.Name, _ = reader.ReadString('\n')
	dep.Name = strings.TrimSpace(dep.Name)

	fmt.Print("Digite os IDs dos projetos separados por vírgula: ")
	projIDsStr, _ := reader.ReadString('\n')
	projIDsStr = strings.TrimSpace(projIDsStr)
	projIDs := strings.Split(projIDsStr, ",")
	for _, idStr := range projIDs {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		projID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("ID de projeto inválido:", idStr)
			continue
		}
		proj, err := ReadProjectByID(projID)
		if err != nil {
			fmt.Println("Projeto não encontrado:", projID)
			continue
		}
		dep.Projects = append(dep.Projects, proj)
	}

	return io.SaveDepartmentToFile(dep)
}

func UpdateDepartment(id int) error {
	departments, err := ReadDepartments()
	if err != nil {
		return fmt.Errorf("Erro ao ler departamentos: %v", err)
	}

	var dept *models.Department
	var index int
	for i, d := range departments {
		if d.ID == id {
			dept = &departments[i]
			index = i
			break
		}
	}

	if dept == nil {
		return fmt.Errorf("Departamento com ID %d não encontrado", id)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		var option int

		fmt.Printf("\nEditando departamento: %d, %s\n", dept.ID, dept.Name)
		fmt.Println("--- Selecione o campo a ser editado ---")
		fmt.Println("1: Nome")
		fmt.Println("2: Gerente (CPF)")
		fmt.Println("0: Salvar e sair")
		fmt.Print("Digite a opção: ")

		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}
		reader.ReadString('\n')

		switch option {
		case 1:
			fmt.Print("Novo nome: ")
			text, _ := reader.ReadString('\n')
			dept.Name = strings.TrimSpace(text)
		case 2:
			fmt.Print("CPF do novo gerente: ")
			cpf, _ := reader.ReadString('\n')
			cpf = strings.TrimSpace(cpf)

			manager, err := ReadEmployeeByCPF(cpf)
			if err != nil {
				fmt.Println("Erro ao buscar gerente:", err)
			} else {
				dept.Manager = manager
				fmt.Println("Gerente atualizado.")
			}
		case 0:
			departments[index] = *dept
			err := io.SaveDepartmentsToFile(departments)
			if err != nil {
				return fmt.Errorf("Erro ao salvar departamento: %v", err)
			}
			fmt.Println("Departamento atualizado com sucesso.")
			return nil
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func DeleteDepartmentByID(id int) error {
	departments, err := ReadDepartments()
	if err != nil {
		return fmt.Errorf("Erro ao ler departamentos: %v", err)
	}

	found := false
	var updated []models.Department
	for _, dept := range departments {
		if dept.ID == id {
			found = true
			continue
		}
		updated = append(updated, dept)
	}

	if !found {
		return fmt.Errorf("Departamento com ID %d não encontrado", id)
	}

	return io.SaveDepartmentsToFile(updated)
}

func ReadDepartmentByID(id int) (*models.Department, error) {
	depts, err := ReadDepartments()
	if err != nil {
		return nil, err
	}

	for _, dept := range depts {
		if dept.ID == id {
			return &dept, nil
		}
	}

	return nil, fmt.Errorf("Departamento com ID %d não encontrado", id)
}

func CreateProject() error {
	var proj models.Project
	reader := bufio.NewReader(os.Stdin)

	proj.ID, _ = getNextProjectID()

	fmt.Print("Digite o nome do projeto: ")
	proj.Name, _ = reader.ReadString('\n')
	proj.Name = strings.TrimSpace(proj.Name)

	fmt.Print("Digite o local do projeto: ")
	proj.Local, _ = reader.ReadString('\n')
	proj.Local = strings.TrimSpace(proj.Local)

	return io.SaveProjectToFile(proj)
}

func UpdateProject(id int) error {
	projects, err := ReadProjects()
	if err != nil {
		return fmt.Errorf("Erro ao ler projetos: %v", err)
	}

	var proj *models.Project
	var index int
	for i, p := range projects {
		if p.ID == id {
			proj = &p
			index = i
			break
		}
	}

	if proj == nil {
		return fmt.Errorf("Projeto com ID %d não encontrado", id)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		var option int

		fmt.Printf("\nEditando projeto: %d, %s, %s\n", proj.ID, proj.Name, proj.Local)
		fmt.Println("--- Selecione o campo a ser editado ---")
		fmt.Println("1: Nome")
		fmt.Println("2: Local")
		fmt.Println("0: Salvar e sair")
		fmt.Print("Digite a opção: ")

		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Erro de entrada. Tente novamente.")
			continue
		}

		reader.ReadString('\n')

		switch option {
		case 1:
			fmt.Print("Novo nome: ")
			text, _ := reader.ReadString('\n')
			proj.Name = strings.TrimSpace(text)
		case 2:
			fmt.Print("Nova local: ")
			text, _ := reader.ReadString('\n')
			proj.Local = strings.TrimSpace(text)
		case 0:
			projects[index] = *proj
			err := io.SaveProjectsToFile(projects)
			if err != nil {
				return fmt.Errorf("erro ao salvar projeto: %v", err)
			}
			fmt.Println("Projeto atualizado com sucesso.")
			return nil
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func ReadProjects() ([]models.Project, error) {
	file, err := os.Open(settings.ProjectFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var projects []models.Project
	scanner := bufio.NewScanner(file)

	var proj models.Project
	lineCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			if lineCount == 3 {
				projects = append(projects, proj)
				proj = models.Project{}
				lineCount = 0
			}
			continue
		}

		if strings.HasPrefix(line, "ID:") {
			idStr := strings.TrimPrefix(line, "ID:")
			idStr = strings.TrimSpace(strings.TrimSuffix(idStr, ","))
			id, err := strconv.Atoi(idStr)
			if err != nil {
				continue
			}
			proj.ID = id
			lineCount++
		} else if strings.HasPrefix(line, "Nome:") {
			proj.Name = strings.TrimSpace(strings.TrimPrefix(line, "Nome:"))
			proj.Name = strings.TrimSuffix(proj.Name, ",")
			lineCount++
		} else if strings.HasPrefix(line, "Local:") {
			proj.Local = strings.TrimSpace(strings.TrimPrefix(line, "Local:"))
			lineCount++
		}
	}

	if lineCount == 3 {
		projects = append(projects, proj)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func DeleteProjectByID(id int) error {
	projects, err := ReadProjects()
	if err != nil {
		return fmt.Errorf("Erro ao ler projetos: %v", err)
	}

	found := false
	var updated []models.Project
	for _, proj := range projects {
		if proj.ID == id {
			found = true
			continue
		}
		updated = append(updated, proj)
	}

	if !found {
		return fmt.Errorf("Projeto com ID %d não encontrado", id)
	}

	return io.SaveProjectsToFile(updated)
}

func ReadProjectByID(id int) (*models.Project, error) {
	projects, err := ReadProjects()
	if err != nil {
		return nil, err
	}

	for _, proj := range projects {
		if proj.ID == id {
			return &proj, nil
		}
	}

	return nil, fmt.Errorf("Projeto com ID %d não encontrado", id)
}

func ReadDepartments() ([]models.Department, error) {
	file, err := os.Open(settings.DepartmentFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var departments []models.Department
	scanner := bufio.NewScanner(file)

	var blockLines []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if len(blockLines) > 0 {
				dept, err := ParseDepartmentBlock(blockLines)
				if err == nil {
					departments = append(departments, dept)
				}
				blockLines = nil
			}
		} else {
			blockLines = append(blockLines, line)
		}
	}

	if len(blockLines) > 0 {
		dept, err := ParseDepartmentBlock(blockLines)
		if err == nil {
			departments = append(departments, dept)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func getNextProjectID() (int, error) {
	lastID, err := common.GetLastID("project")
	if err != nil {
		return 0, err
	}
	newID := lastID + 1
	err = common.UpdateLastID("project", newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func GetNextDepartmentID() (int, error) {
	lastID, err := common.GetLastID("department")
	if err != nil {
		return 0, err
	}
	newID := lastID + 1
	err = common.UpdateLastID("department", newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func SelectManager(id int) error {

	dept, _ := ReadDepartmentByID(id)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o CPF do gerente: ")
	cpf, _ := reader.ReadString('\n')
	cpf = strings.TrimSpace(cpf)

	manager, err := ReadEmployeeByCPF(cpf)
	if err != nil {
		fmt.Println("Nenhum gerente foi encontrado.")
		return err
	}

	if manager.Department.ID == dept.ID {
		dept.Manager = manager
		fmt.Printf("O gerente de CPF %s foi selecionado para o departamento %s.", cpf, dept.Name)
		return nil
	}

	fmt.Printf("O gerente de CPF %s está associado a outro departamento.", cpf)
	return nil
}

func ParseDepartmentBlock(lines []string) (models.Department, error) {
	var dept models.Department

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "ID:") {
			idStr := strings.TrimSpace(strings.TrimPrefix(line, "ID:"))
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return dept, err
			}
			dept.ID = id
		} else if strings.HasPrefix(line, "Nome:") {
			dept.Name = strings.TrimSpace(strings.TrimPrefix(line, "Nome:"))
		} else if strings.HasPrefix(line, "CPF do gerente:") {
			cpf := strings.TrimSpace(strings.TrimPrefix(line, "CPF do gerente:"))
			manager, _ := ReadEmployeeByCPF(cpf)
			dept.Manager = manager
		} else if strings.HasPrefix(line, "ID(s) do(s) projeto(s):") {
			idsStr := strings.TrimSpace(strings.TrimPrefix(line, "ID(s) do(s) projeto(s):"))
			idList := strings.Split(idsStr, ";")
			for _, idStr := range idList {
				idStr = strings.TrimSpace(idStr)
				if idStr == "" {
					continue
				}
				id, err := strconv.Atoi(idStr)
				if err == nil {
					proj, err := ReadProjectByID(id)
					if err == nil {
						dept.Projects = append(dept.Projects, proj)
					}
				}
			}
		}
	}

	return dept, nil
}
