package io

import (
	"empresa/models"
	"empresa/settings"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SaveProjectToFile(proj models.Project) error {
	file, err := os.OpenFile(settings.ProjectFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("ID: %d,\nNome: %s,\nLocal: %s\n\n", proj.ID, proj.Name, proj.Local))
	return err
}

func SaveProjectsToFile(projects []models.Project) error {
	file, err := os.Create(settings.ProjectFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, proj := range projects {
		line := fmt.Sprintf("ID: %d,\nNome: %s,\nLocal: %s\n",
			proj.ID,
			proj.Name,
			proj.Local,
		)

		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func SaveDepartmentToFile(dept models.Department) error {
	file, err := os.OpenFile(settings.DepartmentFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var departmentLine string

	if dept.Manager == nil {
		departmentLine = fmt.Sprintf("ID: %d,\nNome: %s, Nenhum gerente selecionado", dept.ID, dept.Name)
	} else {
		departmentLine = fmt.Sprintf("ID: %d,\nNome: %s,\nCPF do gerente: %s", dept.ID, dept.Name, dept.Manager.CPF)
	}

	projectIDs := []string{}
	for _, proj := range dept.Projects {
		projectIDs = append(projectIDs, strconv.Itoa(proj.ID))
	}

	if len(projectIDs) > 0 {
		departmentLine += ",\nID(s) do(s) projeto(s)" + strings.Join(projectIDs, ";")
	}

	_, err = file.WriteString(departmentLine + "\n")
	return err
}

func SaveDepartmentsToFile(departments []models.Department) error {
	file, err := os.Create(settings.DepartmentFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, dept := range departments {
		var departmentLine string

		if dept.Manager == nil {
			departmentLine = fmt.Sprintf("ID: %d,\nNome: %s,\nNenhum gerente selecionado", dept.ID, dept.Name)
		} else {
			departmentLine = fmt.Sprintf("ID: %d,\nNome: %s,\nCPF do gerente: %s", dept.ID, dept.Name, dept.Manager.CPF)
		}

		projectIDs := []string{}
		for _, proj := range dept.Projects {
			projectIDs = append(projectIDs, strconv.Itoa(proj.ID))
		}

		if len(projectIDs) > 0 {
			departmentLine += ",\nID(s) do(s) projeto(s): " + strings.Join(projectIDs, ";")
		}

		departmentLine += "\n\n"

		_, err := file.WriteString(departmentLine)
		if err != nil {
			return err
		}
	}

	return nil
}

func SaveEmployeeToFile(emp models.Employee) error {
	file, err := os.OpenFile(settings.EmployeeFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	employeeLine :=
		fmt.Sprintf("CPF: %s, Primeiro nome: %s,\nInicial do nome do meio: %s,"+
			"\nÚltimo nome: %s,\nEndereço: %s,\nSalário: %.2f,"+
			"\nGênero: %s,\nData de nascimento: %s",
			emp.FirstName,
			emp.MiddleNameInitial,
			emp.LastName,
			emp.CPF,
			emp.Address,
			emp.Salary,
			emp.Gender,
			emp.DateOfBirth)

	if emp.Department != nil {
		employeeLine += ",\nID do departamento: " + strconv.Itoa(emp.Department.ID)
	} else {
		employeeLine += ","
	}

	projectIDs := []string{}
	for _, proj := range emp.Projects {
		projectIDs = append(projectIDs, strconv.Itoa(proj.ID))
	}

	if len(projectIDs) > 0 {
		employeeLine += ",\nID(s) do(s) projeto(s): " + strings.Join(projectIDs, ";")
	}

	_, err = file.WriteString(employeeLine + "\n\n")
	return err
}

func SaveEmployeesToFile(employees []models.Employee) error {
	file, err := os.Create(settings.EmployeeFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, emp := range employees {
		employeeLine := fmt.Sprintf(
			"CPF: %s,\nPrimeiro nome: %s,\nInicial do nome do meio: %s,\nÚltimo nome: %s,\nEndereço: %s,"+
				"\nSalário: %.2f,\nGênero: %s,\nData de nascimento: %s",
			emp.CPF,
			emp.FirstName,
			emp.MiddleNameInitial,
			emp.LastName,
			emp.Address,
			emp.Salary,
			emp.Gender,
			emp.DateOfBirth,
		)

		if emp.Department != nil {
			employeeLine += ",\nID do departamento: " + strconv.Itoa(emp.Department.ID)
		} else {
			employeeLine += ","
		}

		projectIDs := []string{}
		for _, proj := range emp.Projects {
			projectIDs = append(projectIDs, strconv.Itoa(proj.ID))
		}

		if len(projectIDs) > 0 {
			employeeLine += ",\nID(s) do(s) projeto(s): " + strings.Join(projectIDs, ";")
		}

		employeeLine += "\n\n"

		_, err := file.WriteString(employeeLine)
		if err != nil {
			return err
		}
	}

	return nil
}
