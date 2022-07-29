package adminV1Controller

import (
	"SMT/config"
	"SMT/models"
	"SMT/services"
	adminServices "SMT/services/admin"
	requestTypes "SMT/types/requests"
	responseMessages "SMT/types/responses"

	"github.com/gin-gonic/gin"
)

func AddDepartment(c *gin.Context) {
	var requestBody requestTypes.AddNewDepartment
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		adminServices.SendErrorResponse(c, responseMessages.INVALID_ADD_DEPARTMENT_JSON_INPUT)
	}
	var department models.Department
	department.Code = requestBody.Code
	department.Name = requestBody.Name
	department.Slug = services.CreateSlug(department.Name)
	config.DB.Create(&department)
	adminServices.SuccessResponseWithData(c, responseMessages.DEPARTMENT_ADDED, department)
}

func UpdateDepartment(c *gin.Context) {

}

func DeleteDepartment(c *gin.Context) {

}
