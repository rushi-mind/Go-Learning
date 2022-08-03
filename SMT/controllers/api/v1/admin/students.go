package adminV1Controller

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
	responseTypes "SMT/types/responses"
	"SMT/utility"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	// var students []models.Student
	// config.DB.Find(&students)
	var students []map[string]interface{}
	// config.DB.Raw("select s.id, s.roll_no, s.first_name, s.last_name, s.email, s.semester, s.address, d.id as department_id, d.code as department_code, d.name as department_name from students s left join departments d on s.department_id = d.id").Scan(&students)

	// config.DB.Model(&models.Student{}).Select("students.id, students.first_name, students.last_name, students.email, students.roll_no, students.address, students.semester, departments.code, departments.name").Joins("left join departments on students.department_id = departments.id").Scan(&students)
	config.DB.Model(&models.Student{}).Joins("Department").Find(&students)
	utility.SuccessResponseWithData(c, responseTypes.STUDENTS_FETCHED, students)
}

func GetStudent(c *gin.Context) {
	var studentId int
	studentId, _ = strconv.Atoi(c.Param("id"))
	var student models.Student
	config.DB.First(&student, studentId)
	if student.Id == 0 {
		utility.ErrorResponse(c, responseTypes.INVALID_STUDENT_ID)
		return
	}
	utility.SuccessResponseWithData(c, responseTypes.STUDENT_FETCHED, student)
}

func AddStudent(c *gin.Context) {
	var requestBody requestTypes.AddNewStudent
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	var student models.Student
	student.FirstName = requestBody.FirstName
	student.LastName = requestBody.LastName
	student.DepartmentId = requestBody.DepartmentId
	student.Department = utility.GetDepartment(student.DepartmentId)
	student.Address = requestBody.Address
	student.Semester = requestBody.Semester
	student.RollNo = utility.CreateRollNo(requestBody, student.Department.Code)
	fmt.Println(student.RollNo)
	student.Email = utility.CreateEmailForStudent(student.RollNo)
	student.Password = utility.GetEncryptedPassword(student.RollNo)
	insertResult := config.DB.Create(&student)
	if insertResult.Error != nil {
		utility.ErrorResponse(c, responseTypes.STUDENT_CREATE_ERROR)
		fmt.Println(insertResult.Error)
		return
	}
	utility.SuccessResponseWithData(c, responseTypes.STUDENT_CREATED, student)
}

func UpdateStudent(c *gin.Context) {
	var requestBody requestTypes.UpdateStudent
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, responseTypes.INVALID_STUDENT_ID)
		return
	}
	var count int64
	config.DB.Model(&models.Student{}).Where("id = ?", studentId).Count(&count)
	if count == 0 {
		utility.ErrorResponse(c, responseTypes.INVALID_STUDENT_ID)
		return
	}
	config.DB.Model(&models.Student{}).Updates(requestBody)
	utility.SuccessResponseWithoutData(c, responseTypes.STUDENT_UPDATED)
}

func DeleteStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, responseTypes.INVALID_STUDENT_ID)
		return
	}
	var count int64
	config.DB.Model(&models.Student{}).Where("id = ?", studentId).Count(&count)
	if count == 0 {
		utility.ErrorResponse(c, responseTypes.INVALID_STUDENT_ID)
		return
	}
	config.DB.Delete(&models.Student{}, "id = ?", studentId)
	utility.SuccessResponseWithoutData(c, responseTypes.STUDENT_DELETED)
}
