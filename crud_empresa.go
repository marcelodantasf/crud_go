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
	id    int
	name  string
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
	idTrackerFilename  = "idTracker.txt"
)

func initializeFiles() {
	files := []string{employeeFilename, departmentFilename, projectFilename, idTrackerFilename}

	for _, filename := range files {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}

func getLastID(key string) (int, error) {
	file, err := os.Open(idTrackerFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 && parts[0] == key {
			return strconv.Atoi(parts[1])
		}
	}
	return 0, nil
}

func updateLastID(key string, newID int) error {
	ids := make(map[string]int)

	file, _ := os.Open(idTrackerFilename)
	if file != nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				val, _ := strconv.Atoi(parts[1])
				ids[parts[0]] = val
			}
		}
		file.Close()
	}

	ids[key] = newID

	fileW, err := os.Create(idTrackerFilename)
	if err != nil {
		return err
	}
	defer fileW.Close()

	for k, v := range ids {
		fmt.Fprintf(fileW, "%s:%d\n", k, v)
	}
	return nil
}

func getNextDepartmentID() (int, error) {
	lastID, err := getLastID("department")
	if err != nil {
		return 0, err
	}
	newID := lastID + 1
	err = updateLastID("department", newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func getNextProjectID() (int, error) {
	lastID, err := getLastID("project")
	if err != nil {
		return 0, err
	}
	newID := lastID + 1
	err = updateLastID("project", newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func saveEmployeeToFile(emp Employee) error {
	file, err := os.OpenFile(employeeFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	employeeLine := fmt.Sprintf("%s,%s,%s,%s,%s,%.2f,%s,%s",
		emp.firstName,
		emp.middleNameInitial,
		emp.lastName,
		emp.cpf,
		emp.address,
		emp.salary,
		emp.gender,
		emp.dateOfBirth)

	if emp.department != nil {
		employeeLine += "," + strconv.Itoa(emp.department.id)
	} else {
		employeeLine += ","
	}

	projectIDs := []string{}
	for _, proj := range emp.projects {
		projectIDs = append(projectIDs, strconv.Itoa(proj.id))
	}

	if len(projectIDs) > 0 {
		employeeLine += "," + strings.Join(projectIDs, ";")
	}

	_, err = file.WriteString(employeeLine + "\n")
	return err
}

func saveEmployeesToFile(employees []Employee) error {
	file, err := os.Create(employeeFilename) // sobrescreve o arquivo
	if err != nil {
		return err
	}
	defer file.Close()

	for _, emp := range employees {
		line := fmt.Sprintf("%s,%s,%s,%s,%s,%.2f,%s,%s",
			emp.firstName,
			emp.middleNameInitial,
			emp.lastName,
			emp.cpf,
			emp.address,
			emp.salary,
			emp.gender,
			emp.dateOfBirth)

		if emp.department != nil {
			line += "," + strconv.Itoa(emp.department.id)
		} else {
			line += ","
		}

		projectIDs := []string{}
		for _, proj := range emp.projects {
			projectIDs = append(projectIDs, strconv.Itoa(proj.id))
		}
		if len(projectIDs) > 0 {
			line += "," + strings.Join(projectIDs, ";")
		}

		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func createEmployee() error {
	var emp Employee
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o primeiro nome: ")
	emp.firstName, _ = reader.ReadString('\n')
	emp.firstName = strings.TrimSpace(emp.firstName)

	fmt.Print("Digite a inicial do nome do meio: ")
	emp.middleNameInitial, _ = reader.ReadString('\n')
	emp.middleNameInitial = strings.TrimSpace(emp.middleNameInitial)

	fmt.Print("Digite o último nome: ")
	emp.lastName, _ = reader.ReadString('\n')
	emp.lastName = strings.TrimSpace(emp.lastName)

	fmt.Print("Digite o CPF: ")
	emp.cpf, _ = reader.ReadString('\n')
	emp.cpf = strings.TrimSpace(emp.cpf)

	fmt.Print("Digite o endereço: ")
	emp.address, _ = reader.ReadString('\n')
	emp.address = strings.TrimSpace(emp.address)

	fmt.Print("Digite o salário: ")
	salaryStr, _ := reader.ReadString('\n')
	salaryStr = strings.TrimSpace(salaryStr)
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		return fmt.Errorf("salário inválido")
	}
	emp.salary = salary

	fmt.Print("Digite o gênero (M/F): ")
	emp.gender, _ = reader.ReadString('\n')
	emp.gender = strings.TrimSpace(emp.gender)

	fmt.Print("Data de nascimento (DD/MM/AAAA): ")
	emp.dateOfBirth, _ = reader.ReadString('\n')
	emp.dateOfBirth = strings.TrimSpace(emp.dateOfBirth)

	fmt.Print("Digite o ID do departamento: ")
	deptIDStr, _ := reader.ReadString('\n')
	deptIDStr = strings.TrimSpace(deptIDStr)
	deptID, err := strconv.Atoi(deptIDStr)
	if err != nil {
		return fmt.Errorf("ID do departamento inválido")
	}
	dept, err := readDepartmentByID(deptID)
	if err != nil {
		return err
	}
	emp.department = dept

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
		proj, err := readProjectByID(projID)
		if err != nil {
			fmt.Println("Projeto não encontrado:", projID)
			continue
		}
		emp.projects = append(emp.projects, proj)
	}

	return saveEmployeeToFile(emp)
}

func updateEmployee(cpf string) error {
	employees, err := readEmployees()
	if err != nil {
		return fmt.Errorf("erro ao ler funcionários: %v", err)
	}

	var emp *Employee
	var index int
	for i, e := range employees {
		if e.cpf == cpf {
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

		fmt.Printf("\nEditando funcionário: %s, %s, %s\n", emp.firstName, emp.middleNameInitial, emp.lastName)
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

		// Limpa o buffer de nova linha
		reader.ReadString('\n')

		switch option {
		case 1:
			fmt.Print("Novo primeiro nome: ")
			text, _ := reader.ReadString('\n')
			emp.firstName = strings.TrimSpace(text)
		case 2:
			fmt.Print("Nova inicial do meio: ")
			text, _ := reader.ReadString('\n')
			emp.middleNameInitial = strings.TrimSpace(text)
		case 3:
			fmt.Print("Novo sobrenome: ")
			text, _ := reader.ReadString('\n')
			emp.lastName = strings.TrimSpace(text)
		case 4:
			fmt.Print("Novo endereço: ")
			text, _ := reader.ReadString('\n')
			emp.address = strings.TrimSpace(text)
		case 5:
			fmt.Print("Novo salário: ")
			fmt.Scan(&emp.salary)
		case 6:
			fmt.Print("Novo gênero: ")
			text, _ := reader.ReadString('\n')
			emp.gender = strings.TrimSpace(text)
		case 7:
			fmt.Print("Nova data de nascimento: ")
			text, _ := reader.ReadString('\n')
			emp.dateOfBirth = strings.TrimSpace(text)
		case 8:
			fmt.Print("Digite o novo ID do departamento: ")
			deptIDStr, _ := reader.ReadString('\n')
			deptIDStr = strings.TrimSpace(deptIDStr)
			deptID, err := strconv.Atoi(deptIDStr)
			if err != nil {
				return fmt.Errorf("ID do departamento inválido")
			}
			dept, err := readDepartmentByID(deptID)
			if err != nil {
				return err
			}
			emp.department = dept
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
				proj, err := readProjectByID(projID)
				if err != nil {
					fmt.Println("Projeto não encontrado:", projID)
					continue
				}
				emp.projects = append(emp.projects, proj)
			}
		case 0:
			employees[index] = *emp
			err := saveEmployeesToFile(employees)
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

func deleteEmployeeByCPF(cpf string) error {
	employees, err := readEmployees()
	if err != nil {
		return fmt.Errorf("Erro ao ler funcionários: %v", err)
	}

	found := false
	var updated []Employee
	for _, emp := range employees {
		if emp.cpf == cpf {
			found = true
			continue
		}
		updated = append(updated, emp)
	}

	if !found {
		return fmt.Errorf("Funcionário com CPF %s não encontrado", cpf)
	}

	return saveEmployeesToFile(updated)
}

// Falta add a exception

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

	for i := range employees {
		if employees[i].cpf == cpf {
			return &employees[i], nil
		}
	}

	return nil, fmt.Errorf("Funcionário com CPF %s não encontrado", cpf)
}

func saveProjectToFile(proj Project) error {
	file, err := os.OpenFile(projectFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d,%s,%s\n", proj.id, proj.name, proj.local))
	return err
}

func saveProjectsToFile(projects []Project) error {
	file, err := os.Create(projectFilename) // sobrescreve o arquivo
	if err != nil {
		return err
	}
	defer file.Close()

	for _, proj := range projects {
		line := fmt.Sprintf("%d,%s,%s",
			proj.id,
			proj.name,
			proj.local,
		)

		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func createProject() error {
	var proj Project
	reader := bufio.NewReader(os.Stdin)

	proj.id, _ = getNextProjectID()

	fmt.Print("Digite o nome do projeto: ")
	proj.name, _ = reader.ReadString('\n')
	proj.name = strings.TrimSpace(proj.name)

	fmt.Print("Digite o local do projeto: ")
	proj.local, _ = reader.ReadString('\n')
	proj.local = strings.TrimSpace(proj.local)

	return saveProjectToFile(proj)
}

func updateProject(id int) error {
	projects, err := readProjects()
	if err != nil {
		return fmt.Errorf("Erro ao ler projetos: %v", err)
	}

	var proj *Project
	var index int
	for i, p := range projects {
		if p.id == id {
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

		fmt.Printf("\nEditando projeto: %d, %s, %s\n", proj.id, proj.name, proj.local)
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
			proj.name = strings.TrimSpace(text)
		case 2:
			fmt.Print("Nova local: ")
			text, _ := reader.ReadString('\n')
			proj.local = strings.TrimSpace(text)
		case 0:
			projects[index] = *proj
			err := saveProjectsToFile(projects)
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

func deleteProjectByID(id int) error {
	projects, err := readProjects()
	if err != nil {
		return fmt.Errorf("Erro ao ler projetos: %v", err)
	}

	found := false
	var updated []Project
	for _, proj := range projects {
		if proj.id == id {
			found = true
			continue
		}
		updated = append(updated, proj)
	}

	if !found {
		return fmt.Errorf("Projeto com ID %d não encontrado", id)
	}

	return saveProjectsToFile(updated)
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

func saveDepartmentToFile(dept Department) error {
	file, err := os.OpenFile(departmentFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	departmentLine := fmt.Sprintf("%d,%s, Nenhum gerente selecionado", dept.id, dept.name)

	projectIDs := []string{}
	for _, proj := range dept.projects {
		projectIDs = append(projectIDs, strconv.Itoa(proj.id))
	}

	if len(projectIDs) > 0 {
		departmentLine += "," + strings.Join(projectIDs, ";")
	}

	_, err = file.WriteString(departmentLine + "\n")
	return err
}

func saveDepartmentsToFile(departments []Department) error {
	file, err := os.Create(departmentFilename) // sobrescreve o arquivo
	if err != nil {
		return err
	}
	defer file.Close()

	for _, dept := range departments {
		line := fmt.Sprintf("%d,%s", dept.id, dept.name)

		if dept.manager != nil {
			line += "," + dept.manager.cpf
		} else {
			line += ","
		}

		projectIDs := []string{}
		for _, proj := range dept.projects {
			projectIDs = append(projectIDs, strconv.Itoa(proj.id))
		}
		if len(projectIDs) > 0 {
			line += "," + strings.Join(projectIDs, ";")
		}

		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func createDepartment() error {
	var dep Department
	reader := bufio.NewReader(os.Stdin)

	dep.id, _ = getNextDepartmentID()

	fmt.Print("Digite o nome do departamento: ")
	dep.name, _ = reader.ReadString('\n')
	dep.name = strings.TrimSpace(dep.name)

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
		proj, err := readProjectByID(projID)
		if err != nil {
			fmt.Println("Projeto não encontrado:", projID)
			continue
		}
		dep.projects = append(dep.projects, proj)
	}

	return saveDepartmentToFile(dep)
}

func selectManager(id int) error {

	dept, _ := readDepartmentByID(id)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o CPF do gerente: ")
	cpf, _ := reader.ReadString('\n')
	cpf = strings.TrimSpace(cpf)

	manager, err := readEmployeeByCPF(cpf)
	if err != nil {
		fmt.Println("Nenhum gerente foi encontrado.")
		return err
	}

	if manager.department.id == dept.id {
		dept.manager = manager
		fmt.Printf("O gerente de CPF %s foi selecionado para o departamento %s.", cpf, dept.name)
		return nil
	}

	fmt.Printf("O gerente de CPF %s está associado a outro departamento.", cpf)
	return nil
}

func updateDepartment(id int) error {
	departments, err := readDepartments()
	if err != nil {
		return fmt.Errorf("Erro ao ler departamentos: %v", err)
	}

	var dept *Department
	var index int
	for i, d := range departments {
		if d.id == id {
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

		fmt.Printf("\nEditando departamento: %d, %s\n", dept.id, dept.name)
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
			dept.name = strings.TrimSpace(text)
		case 2:
			fmt.Print("CPF do novo gerente: ")
			cpf, _ := reader.ReadString('\n')
			cpf = strings.TrimSpace(cpf)

			manager, err := readEmployeeByCPF(cpf)
			if err != nil {
				fmt.Println("Erro ao buscar gerente:", err)
			} else {
				dept.manager = manager
				fmt.Println("Gerente atualizado.")
			}
		case 0:
			departments[index] = *dept
			err := saveDepartmentsToFile(departments)
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

func deleteDepartmentByID(id int) error {
	departments, err := readDepartments()
	if err != nil {
		return fmt.Errorf("Erro ao ler departamentos: %v", err)
	}

	found := false
	var updated []Department
	for _, dept := range departments {
		if dept.id == id {
			found = true
			continue
		}
		updated = append(updated, dept)
	}

	if !found {
		return fmt.Errorf("Departamento com ID %d não encontrado", id)
	}

	return saveDepartmentsToFile(updated)
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
			createEmployee()
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
			fmt.Print("Digite o CPF do funcionário a ser alterado: ")
			var cpf string
			fmt.Scan(&cpf)
			updateEmployee(cpf)

			//fmt.Println("Alterando funcionário...")
		case 4:
			fmt.Println("Excluindo funcionário...")
			fmt.Print("Digite o CPF do funcionário a ser deletado: ")
			var cpf string
			fmt.Scan(&cpf)
			deleteEmployeeByCPF(cpf)
		case 5:
			fmt.Println("Buscar funcionário por CPF...")
			fmt.Print("Digite o CPF: ")
			var cpf string
			fmt.Scan(&cpf)

			reader.ReadString('\n')

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
			createProject()
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
			fmt.Print("Digite o ID do projeto a ser alterado: ")
			var id int
			fmt.Scan(&id)
			updateProject(id)
		case 4:
			fmt.Println("Excluindo projeto...")
			fmt.Print("Digite o ID do projeto a ser deletado: ")
			var id int
			fmt.Scan(&id)
			deleteProjectByID(id)
		case 5:
			fmt.Println("Buscar projeto por ID...")
			fmt.Print("Digite o ID do projeto: ")
			var id int
			fmt.Scan(&id)

			reader.ReadString('\n')

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
			err := createDepartment()
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
			updateDepartment(id)
		case 4:
			fmt.Println("Excluindo departamento...")
			fmt.Print("Digite o ID do departamento a ser alterado: ")
			var id int
			fmt.Scan(&id)
			deleteDepartmentByID(id)
		case 5:
			fmt.Println("Buscando departamento...")
			fmt.Print("Digite o ID do departamento: ")
			var id int
			fmt.Scan(&id)

			reader.ReadString('\n')

			dept, err := readDepartmentByID(id)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("ID: %d | Nome: %s\n", dept.id, dept.name)
			}
		case 6:
			fmt.Print("Digite o ID do departamento para definir o gerente: ")
			var id int
			fmt.Scan(&id)
			selectManager(id)
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
