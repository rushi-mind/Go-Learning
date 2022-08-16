package repository

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
	"strconv"
)

func AddAssignment(requestBody requestTypes.AddUpdateAssignment) error {
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

func UpdateAssignment(id int, requestBody requestTypes.AddUpdateAssignment) error {
	query := "UPDATE assignments SET "
	if requestBody.Name != "" {
		query += "name = '" + requestBody.Name + "', "
	}
	if requestBody.Semester != "" {
		query += "semester = '" + requestBody.Semester + "', "
	}
	if requestBody.DepartmentID != 0 {
		query += "department_id = " + strconv.Itoa(int(requestBody.DepartmentID)) + ", "
	}
	if requestBody.Deadline != "" {
		query += "deadline = '" + requestBody.Deadline + "', "
	}
	query = string([]byte(query)[:len(query)-2])
	query += "WHERE id = " + strconv.Itoa(id)
	return config.DB.Exec(query).Error
}

func IsValidAssignmentID(id int) bool {
	var count int
	config.DB.Raw("SELECT COUNT(id) FROM assignments WHERE id = ?", id).Scan(&count)
	return count != 0
}
