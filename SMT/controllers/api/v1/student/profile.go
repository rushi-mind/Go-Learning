package studentV1Controllers

import (
	"SMT/repository"
	requestTypes "SMT/types/requests"
	stringTypes "SMT/types/strings"
	"SMT/utility"

	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	var requestBody requestTypes.ChangePassword
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	payload, _ := c.Get("user")
	id := uint((payload.(map[string]interface{})["id"]).(float64))
	encryptedPassword := repository.GetPasswordFromDB(id)
	if !utility.ValidatePassword(requestBody.OldPassword, encryptedPassword) {
		utility.ErrorResponse(c, stringTypes.INVALID_PASSWORD)
		return
	}
	if !repository.UpdateStudentPassword(id, requestBody.NewPassword) {
		utility.ErrorResponse(c, stringTypes.PASSWORD_CHANGE_FAILURE)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.PASSWORD_CHANGED)
}

func GetProfile(c *gin.Context) {
	payload, _ := c.Get("user")
	id := int((payload.(map[string]interface{})["id"]).(float64))
	student := repository.GetStudentByID(id)
	utility.SuccessResponseWithData(c, stringTypes.PROFILE_FETCHED, student, 1)
}
