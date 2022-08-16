package repository

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
	responseTypes "SMT/types/responses"
	"SMT/utility"
	"fmt"
	"strconv"
)

func IsValidDepartmentID(id int) bool {
	var count int
	config.DB.Raw("SELECT COUNT(1) FROM departments WHERE id = ?", id).Scan(&count)
	return count != 0
}

func AddDepartment(department models.Department) error {
	return config.DB.Exec("INSERT INTO departments(code, name, slug) VALUES(?, ?, ?)", department.Code, department.Name, department.Slug).Error
}

func UpdateDepartment(department requestTypes.UpdateDepartment, id int) error {
	query := "UPDATE departments SET "
	if department.Name != "" {
		query += "name = '" + department.Name + "', "
		query += "slug = '" + utility.CreateSlug(department.Name) + "'"
	}
	if department.Code != "" {
		if department.Name != "" {
			query += ", "
		}
		query += "code = '" + department.Code + "'"
	}
	query += " WHERE id = " + strconv.Itoa(id)
	fmt.Println("query: ", query)
	return config.DB.Exec(query).Error
}

func GetDepartments() []responseTypes.GetDepartment {
	var departments []responseTypes.GetDepartment
	config.DB.Raw("SELECT id, name, code FROM departments").Scan(&departments)
	return departments
}

func GetDepartment(id int) responseTypes.GetDepartment {
	var department responseTypes.GetDepartment
	config.DB.Raw("SELECT id, name, code FROM departments WHERE id = ??", id).Scan(&department)
	return department
}
