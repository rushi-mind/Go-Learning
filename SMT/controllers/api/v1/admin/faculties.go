package adminV1Controller

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
	responseTypes "SMT/types/responses"
	"SMT/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTeachers(c *gin.Context) {
	var teachers []models.Faculty
	config.DB.Find(&teachers)
	utility.SuccessResponseWithData(c, responseTypes.TEACHERS_FETCHED, teachers)
}

func GetTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, responseTypes.INVALID_TEACHER_ID)
		return
	}
	var teacher models.Faculty
	config.DB.First(&teacher, id)
	if teacher.Id == 0 {
		utility.ErrorResponse(c, responseTypes.INVALID_TEACHER_ID)
		return
	}
	utility.SuccessResponseWithData(c, responseTypes.TEACHER_FETCHED, teacher)
}

func AddTeacher(c *gin.Context) {
	var requestBody requestTypes.AddNewTeacher
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	var teacher models.Faculty
	teacher.FirstName = requestBody.FirstName
	teacher.LastName = requestBody.LastName
	teacher.DepartmentId = requestBody.DepartmentId
	insertResult := config.DB.Create(&teacher)
	if insertResult.Error != nil {
		utility.ErrorResponse(c, responseTypes.TEACHER_CREATE_ERROR)
		return
	}
	utility.SuccessResponseWithData(c, responseTypes.TEACHER_CREATED, teacher)
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
		utility.ErrorResponse(c, responseTypes.INVALID_TEACHER_ID)
		return
	}
	var count int64
	config.DB.Model(&models.Faculty{}).Where("id = ?", teacherId).Count(&count)
	if count == 0 {
		utility.ErrorResponse(c, responseTypes.INVALID_TEACHER_ID)
		return
	}
	config.DB.Model(&models.Faculty{}).Updates(requestBody)
	utility.SuccessResponseWithoutData(c, responseTypes.TEACHER_UPDATE)
}

func DeleteTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, responseTypes.INVALID_TEACHER_ID)
		return
	}
	var count int64
	config.DB.Model(&models.Faculty{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		utility.ErrorResponse(c, responseTypes.INVALID_TEACHER_ID)
		return
	}
	config.DB.Delete(&models.Faculty{}).Where("id = ?", id)
	utility.SuccessResponseWithoutData(c, responseTypes.TEACHER_DELETED)
}
