package repository

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
)

func AddAssignment(requestBody requestTypes.AddAssignment) error {
	return config.DB.Exec("INSERT INTO assignments(name, semester, department_id, deadline) VALUES(?, ?, ?, ?)", requestBody.Name, requestBody.Semester, requestBody.DepartmentID, requestBody.Deadline).Error
}

func DeleteAssignment(id int) bool {
	return config.DB.Exec("DELETE FROM assignments WHERE id = ?", id).RowsAffected != 0
}

func GetAssignments(deptID string, semester string) []models.Assignment {
	var assignments []models.Assignment
	config.DB.Raw("SELECT * FROM assignments WHERE semester = ? AND department_id = ?", semester, deptID).Scan(&assignments)
	return assignments
}
