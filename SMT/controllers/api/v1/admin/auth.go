package adminV1Controller

import (
	"SMT/config"
	"SMT/models"
	requestTypes "SMT/types/requests"
	stringTypes "SMT/types/strings"
	"SMT/utility"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var requestBody requestTypes.AdminAuth
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}

	var admin models.Admin
	getAdminResult := config.DB.First(&admin, "admin_id = ?", requestBody.AdminId)
	if getAdminResult.RowsAffected == 0 {
		utility.ErrorResponse(c, stringTypes.INVALID_ADMINID)
		return
	}
	if !utility.ValidatePassword(requestBody.Password, admin.Password) {
		utility.ErrorResponse(c, stringTypes.INVALID_PASSWORD)
		return
	}

	token, err := utility.CreateJWT(map[string]interface{}{
		"id":       admin.Id,
		"admin_id": admin.AdminId,
		"email":    admin.EmailId,
		"role":     "admin",
	})
	if err != nil {
		utility.ErrorResponse(c, stringTypes.LOGIN_FAILED)
		return
	}

	utility.SuccessResponseWithToken(c, token)
}

func ChangePassword(c *gin.Context) {
	var requestBody requestTypes.ChangePassword
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	jwtPayload, _ := c.Get("user")
	adminId := jwtPayload.(map[string]interface{})["admin_id"]
	var admin models.Admin
	config.DB.First(&admin, "admin_id = ?", adminId)
	if !utility.ValidatePassword(requestBody.OldPassword, admin.Password) {
		utility.ErrorResponse(c, stringTypes.INVALID_PASSWORD)
		return
	}
	admin.Password = utility.GetEncryptedPassword(requestBody.NewPassword)
	config.DB.Save(&admin)
	utility.SuccessResponseWithoutData(c, stringTypes.PASSWORD_CHANGED)
}
