package adminV1Controller

import (
	"SMT/config"
	"SMT/models"
	"SMT/repository"
	requestTypes "SMT/types/requests"
	stringTypes "SMT/types/strings"
	"SMT/utility"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	students := repository.GetAllStudents()
	utility.SuccessResponseWithData(c, stringTypes.STUDENTS_FETCHED, students)
}

func GetStudent(c *gin.Context) {
	var studentId int
	studentId, _ = strconv.Atoi(c.Param("id"))
	student := repository.GetStudentByID(studentId)
	if student.ID == 0 {
		utility.ErrorResponse(c, stringTypes.INVALID_STUDENT_ID)
		return
	}
	utility.SuccessResponseWithData(c, stringTypes.STUDENT_FETCHED, student)
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
	student.Email = utility.CreateEmailForStudent(student.RollNo)
	student.Password = utility.GetEncryptedPassword(student.RollNo)
	if err := repository.InsertStudent(student); err != nil {
		log.Default().Println(err)
		utility.ErrorResponse(c, stringTypes.STUDENT_CREATE_ERROR)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.STUDENT_CREATED)
}

func UpdateStudent(c *gin.Context) {
	var requestBody requestTypes.UpdateStudent
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, stringTypes.INVALID_STUDENT_ID)
		return
	}
	if !repository.IsValidStudentId(studentId) {
		utility.ErrorResponse(c, stringTypes.INVALID_STUDENT_ID)
		return
	}
	if err := repository.UpdateStudent(requestBody, studentId); err != nil {
		log.Default().Println(err)
		utility.ErrorResponse(c, stringTypes.STUDENT_UPDATE_FAIL)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.STUDENT_UPDATED)
}

func DeleteStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Default().Println(err)
		utility.ErrorResponse(c, stringTypes.INVALID_STUDENT_ID)
		return
	}
	if !repository.IsValidStudentId(studentId) {
		utility.ErrorResponse(c, stringTypes.INVALID_STUDENT_ID)
		return
	}
	config.DB.Delete(&models.Student{}, "id = ?", studentId)
	utility.SuccessResponseWithoutData(c, stringTypes.STUDENT_DELETED)
}
