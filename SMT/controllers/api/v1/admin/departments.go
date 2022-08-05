package adminV1Controller

import (
	"SMT/config"
	"SMT/models"
	"SMT/repository"
	requestTypes "SMT/types/requests"
	stringTypes "SMT/types/strings"
	"SMT/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddDepartment(c *gin.Context) {
	var requestBody requestTypes.AddNewDepartment
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	var department models.Department
	department.Code = requestBody.Code
	department.Name = requestBody.Name
	department.Slug = utility.CreateSlug(department.Name)
	if err := repository.AddDepartment(department); err != nil {
		utility.ErrorResponse(c, stringTypes.DEPARTMENT_INSERT_FAIL)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.DEPARTMENT_ADDED)
}

func UpdateDepartment(c *gin.Context) {
	var requestBody requestTypes.UpdateDepartment
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	id, err := strconv.Atoi(c.Param("deptId"))
	if err != nil {
		utility.ErrorResponse(c, stringTypes.INVALID_DEPARTMENT_ID)
		return
	}
	if !repository.IsValidDepartmentID(id) {
		utility.ErrorResponse(c, stringTypes.DEPARTMENT_NOT_FOUND)
		return
	}
	if err := repository.UpdateDepartment(requestBody, id); err != nil {
		utility.ErrorResponse(c, stringTypes.DEPARTMENT_UPDATE_FAIL)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.DEPARTMENT_UPDATED)
}

func DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("deptId"))
	if err != nil {
		utility.ErrorResponse(c, stringTypes.INVALID_DEPARTMENT_ID)
		return
	}
	var department models.Department
	config.DB.First(&department, id)
	if department.Id == 0 {
		utility.ErrorResponse(c, stringTypes.DEPARTMENT_NOT_FOUND)
		return
	}
	config.DB.Delete(&department)
	utility.SuccessResponseWithoutData(c, stringTypes.DEPARTMENT_DELETED)
}
