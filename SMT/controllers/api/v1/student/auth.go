package studentV1Controllers

import (
	"SMT/repository"
	requestTypes "SMT/types/requests"
	stringTypes "SMT/types/strings"
	"SMT/utility"

	"github.com/gin-gonic/gin"
)

func AuthStudent(c *gin.Context) {
	var requestBody requestTypes.AuthStudent
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utility.ErrorResponse(c, err.Error())
		return
	}
	var studentID uint
	var isValid bool
	if studentID, isValid = repository.IsValidRollNo(requestBody.RollNo); !isValid {
		utility.ErrorResponse(c, stringTypes.INVALID_ROLL_NO)
		return
	}

	encryptedPassword := repository.GetPasswordFromDB(studentID)
	if !utility.ValidatePassword(requestBody.Password, encryptedPassword) {
		utility.ErrorResponse(c, stringTypes.INVALID_PASSWORD)
		return
	}

	token, err := utility.CreateJWT(map[string]interface{}{
		"id":      studentID,
		"roll_no": requestBody.RollNo,
		"role":    "student",
	})
	if err != nil {
		utility.ErrorResponse(c, stringTypes.AUTHENTICATION_FAILED)
		return
	}
	utility.SuccessResponseWithToken(c, token)
}
