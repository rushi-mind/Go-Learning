package repository

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
	responseTypes "SMT/types/responses"
	"strconv"
)

func IsValidStudentId(studentId int) bool {
	var count int
	config.DB.Raw("SELECT count(1) FROM students WHERE id = ?", studentId).Scan(&count)
	return count != 0
}

func InsertStudent(student models.Student) error {
	return config.DB.Exec("INSERT INTO students(first_name, last_name, department_id, roll_no, semester, email, password, address) VALUES(?, ?, ?, ?, ?, ?, ?, ?);", student.FirstName, student.LastName, student.DepartmentId, student.RollNo, student.Semester, student.Email, student.Password, student.Address).Error
}

func GetStudentByID(id int) responseTypes.StudentResponse {
	var student responseTypes.StudentResponse
	config.DB.Raw("SELECT s.id, roll_no, first_name, last_name, email, semester, address, profile_image, department_id, d.code as department_code, d.name as department_name FROM students s LEFT JOIN departments d ON department_id = d.id WHERE s.id = ?", id).Scan(&student)
	return student
}

func GetAllStudents() []responseTypes.StudentResponse {
	var students []responseTypes.StudentResponse
	config.DB.Raw("SELECT s.id, roll_no, first_name, last_name, email, semester, address, profile_image, department_id, d.code as department_code, d.name as department_name FROM students s LEFT JOIN departments d ON department_id = d.id").Scan(&students)
	return students
}

func UpdateStudent(requestBody requestTypes.UpdateStudent, id int) error {
	query := "UPDATE students SET "
	if requestBody.FirstName != "" {
		query += "first_name = \"" + requestBody.FirstName + "\", "
	}
	if requestBody.LastName != "" {
		query += "last_name = \"" + requestBody.LastName + "\", "
	}
	if requestBody.Address != "" {
		query += "address = \"" + requestBody.Address + "\", "
	}
	query += "\b\b WHERE id = " + strconv.Itoa(id) + ";"
	return config.DB.Raw(query).Error
}

func DeleteStudent(id int) {
	config.DB.Exec("DELETE FROM students WHERE id = ?", id)
}

func GetStudentsListOfClass(deptID, semester string) []responseTypes.StudentResponse {
	var students []responseTypes.StudentResponse
	config.DB.Raw("SELECT s.id, roll_no, first_name, last_name, email, semester, address, profile_image, department_id, d.code as department_code, d.name as department_name FROM students s LEFT JOIN departments d ON department_id = d.id WHERE s.department_id = ? AND s.semester = ?", deptID, semester).Scan(&students)
	return students
}

func GetStudentByRollNo(rollNo string) responseTypes.StudentResponse {
	var student responseTypes.StudentResponse
	config.DB.Raw("SELECT s.id, roll_no, first_name, last_name, email, semester, address, profile_image, department_id, d.code as department_code, d.name as department_name FROM students s LEFT JOIN departments d ON department_id = d.id WHERE roll_no = ?", rollNo).Scan(&student)
	return student
}

func GetPasswordFromDB(studentID uint) string {
	var password string
	config.DB.Raw("SELECT password FROM students WHERE id = ?", studentID).Scan(&password)
	return password
}

func IsValidRollNo(rollNo string) (uint, bool) {
	var id uint
	config.DB.Raw("SELECT id from students WHERE roll_no = ?", rollNo).Scan(&id)
	return id, id != 0
}

func UpdateStudentPassword(studentID uint, newPassword string) bool {
	return config.DB.Exec("UPDATE students SET password = ? WHERE id = ?", newPassword, studentID).RowsAffected > 0
}
