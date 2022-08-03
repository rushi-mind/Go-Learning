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

func AddDepartment(c *gin.Context) {
	var requestBody requestTypes.AddNewDepartment
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	var department models.Department
	department.Code = requestBody.Code
	department.Name = requestBody.Name
	department.Slug = utility.CreateSlug(department.Name)
	config.DB.Create(&department)
	utility.SuccessResponseWithData(c, responseTypes.DEPARTMENT_ADDED, department)
}

func UpdateDepartment(c *gin.Context) {
	var requestBody requestTypes.UpdateDepartment
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	id, err := strconv.Atoi(c.Param("deptId"))
	if err != nil {
		utility.ErrorResponse(c, responseTypes.INVALID_DEPARTMENT_ID)
		return
	}
	var count int64
	config.DB.Model(&models.Department{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		utility.ErrorResponse(c, responseTypes.DEPARTMENT_NOT_FOUND)
		return
	}
	config.DB.Model(&models.Department{}).Updates(requestBody)
	utility.SuccessResponseWithoutData(c, responseTypes.DEPARTMENT_UPDATED)
}

func DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("deptId"))
	if err != nil {
		utility.ErrorResponse(c, responseTypes.INVALID_DEPARTMENT_ID)
		return
	}
	var department models.Department
	config.DB.First(&department, id)
	if department.Id == 0 {
		utility.ErrorResponse(c, responseTypes.DEPARTMENT_NOT_FOUND)
		return
	}
	config.DB.Delete(&department)
	utility.SuccessResponseWithoutData(c, responseTypes.DEPARTMENT_DELETED)
}
