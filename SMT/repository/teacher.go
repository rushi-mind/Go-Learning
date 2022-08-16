package repository

import (
	"SMT/config"
	responseTypes "SMT/types/responses"
)

func IsValidEmployeeID(empID string) bool {
	var count int64
	config.DB.Raw("SELECT COUNT(id) FROM faculties WHERE employee_id = ?", empID).Count(&count)
	return count != 0
}

func GetTeacherByEmployeeID(empID string) responseTypes.TeacherResponse {
	var teacher responseTypes.TeacherResponse
	config.DB.Raw("SELECT f.id, f.first_name, f.last_name, f.employee_id, f.email, f.department_id, d.code as department_code, d.name as department_name FROM faculties f LEFT JOIN departments d ON f.department_id = d.id WHERE employee_id = ?", empID).Scan(&teacher)
	return teacher
}

func GetTeacherPasswordFromDB(empID string) string {
	var password string
	config.DB.Raw("SELECT password FROM faculties WHERE employee_id = ?", empID).Scan(&password)
	return password
}

func UpdateTeacherPassword(id uint, password string) error {
	return config.DB.Exec("UPDATE faculties SET password = ? WHERE id = ?", password, id).Error
}
