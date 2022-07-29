package adminV1Controller

import (
	"SMT/config"
	"SMT/models"
	adminServices "SMT/services/admin"
	requestTypes "SMT/types/requests"
	responseMessages "SMT/types/responses"
	"SMT/utility"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var requestBody requestTypes.AdminAuth
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		adminServices.SendErrorResponse(c, responseMessages.ADMIN_AUTH_INVALID_JSON_INPUT)
		return
	}

	var admin models.Admin
	getAdminResult := config.DB.First(&admin, "adminId = ?", requestBody.AdminId)
	if getAdminResult.RowsAffected == 0 {
		adminServices.SendErrorResponse(c, responseMessages.INVALID_ADMINID)
		return
	}
	if !utility.ValidatePassword(requestBody.Password, admin.Password) {
		adminServices.SendErrorResponse(c, responseMessages.INVALID_PASSWORD)
		return
	}

	token, err := utility.CreateJWT(map[string]interface{}{
		"id":      admin.Id,
		"adminId": admin.AdminId,
		"email":   admin.EmailId,
	})
	if err != nil {
		adminServices.SendErrorResponse(c, responseMessages.LOGIN_FAILED)
		return
	}

	adminServices.SendLoginSuccessResponse(c, token)
}

func ChangePassword(c *gin.Context) {
	var requestBody requestTypes.AdminChangePassword
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		adminServices.SendErrorResponse(c, responseMessages.INVALID_ADMIN_CHANGE_PASSWORD_JSON_INPUT)
		return
	}
	jwtPayload, _ := c.Get("user")
	adminId := jwtPayload.(map[string]interface{})["adminId"]
	var admin models.Admin
	config.DB.First(&admin, "adminId = ?", adminId)
	if !utility.ValidatePassword(requestBody.OldPassword, admin.Password) {
		adminServices.SendErrorResponse(c, responseMessages.INVALID_PASSWORD)
		return
	}
	admin.Password = utility.GetEncryptedPassword(requestBody.NewPassword)
	config.DB.Save(&admin)
	adminServices.SendSuccessResponseWithoutBody(c, responseMessages.PASSWORD_CHANGED)
}
