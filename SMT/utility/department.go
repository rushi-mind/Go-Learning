package utility

import (
	"SMT/config"
	"SMT/models"
	"strings"
)

func GetDepartment(deptId uint) models.Department {
	var department models.Department
	config.DB.First(&department, deptId)
	return department
}

func CreateSlug(str string) string {
	return strings.ToLower(strings.Join(strings.Split(str, " "), "-"))
}
