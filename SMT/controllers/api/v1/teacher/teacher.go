package teacherV1Controller

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
	stringTypes "SMT/types/strings"
	"SMT/utility"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTeachers(c *gin.Context) {
	var teachers []models.Faculty
	config.DB.Find(&teachers)
	utility.SuccessResponseWithData(c, stringTypes.TEACHERS_FETCHED, teachers, len(teachers))
}

func GetTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, stringTypes.INVALID_TEACHER_ID)
		return
	}
	var teacher models.Faculty
	config.DB.First(&teacher, id)
	if teacher.Id == 0 {
		utility.ErrorResponse(c, stringTypes.INVALID_TEACHER_ID)
		return
	}
	utility.SuccessResponseWithData(c, stringTypes.TEACHER_FETCHED, teacher, 1)
}

func AddTeacher(c *gin.Context) {
	var requestBody requestTypes.AddNewTeacher
	if c.Request.Body == nil {
		fmt.Println("invalid here")
		utility.ErrorResponse(c, stringTypes.INVALID_INPUT_JSON)
		return
	}
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	var teacher models.Faculty
	teacher.FirstName = requestBody.FirstName
	teacher.LastName = requestBody.LastName
	teacher.DepartmentId = requestBody.DepartmentId
	teacher.EmployeeId = utility.CreateEmployeeId(teacher.DepartmentId)
	if teacher.EmployeeId == "" {
		utility.ErrorResponse(c, stringTypes.INVALID_DEPARTMENT_ID)
		return
	}
	insertResult := config.DB.Create(&teacher)
	if insertResult.Error != nil {
		utility.ErrorResponse(c, stringTypes.TEACHER_CREATE_ERROR)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.TEACHER_CREATED)
}

func UpdateTeacher(c *gin.Context) {
	var requestBody requestTypes.UpdateTeacher
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	teacherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, stringTypes.INVALID_TEACHER_ID)
		return
	}
	var count int64
	config.DB.Model(&models.Faculty{}).Where("id = ?", teacherId).Count(&count)
	if count == 0 {
		utility.ErrorResponse(c, stringTypes.INVALID_TEACHER_ID)
		return
	}
	config.DB.Model(&models.Faculty{}).Updates(requestBody)
	utility.SuccessResponseWithoutData(c, stringTypes.TEACHER_UPDATE)
}

func DeleteTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, stringTypes.INVALID_TEACHER_ID)
		return
	}
	var count int64
	config.DB.Model(&models.Faculty{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		utility.ErrorResponse(c, stringTypes.INVALID_TEACHER_ID)
		return
	}
	config.DB.Delete(&models.Faculty{}).Where("id = ?", id)
	utility.SuccessResponseWithoutData(c, stringTypes.TEACHER_DELETED)
}
