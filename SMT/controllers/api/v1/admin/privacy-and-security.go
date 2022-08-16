package adminV1Controller

import (
	"SMT/config"
	"SMT/models"
	"SMT/repository"
	requestTypes "SMT/types/requests"
	stringTypes "SMT/types/strings"
	"SMT/utility"

	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	var requestBody requestTypes.ChangePassword
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	jwtPayload, _ := c.Get("user")
	switch jwtPayload.(map[string]interface{})["role"] {
	case "admin":
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
		break
	case "teacher":
		encryptedPassword := repository.GetTeacherPasswordFromDB((jwtPayload.(map[string]interface{})["employee_id"]).(string))
		if !utility.ValidatePassword(requestBody.OldPassword, encryptedPassword) {
			utility.ErrorResponse(c, stringTypes.INVALID_PASSWORD)
			return
		}
		if err := repository.UpdateTeacherPassword((jwtPayload.(map[string]interface{})["id"]).(uint), utility.GetEncryptedPassword(requestBody.NewPassword)); err != nil {
			utility.ErrorResponse(c, stringTypes.PASSWORD_CHANGE_FAILURE)
			return
		}
		utility.SuccessResponseWithoutData(c, stringTypes.PASSWORD_CHANGED)
	}

}
