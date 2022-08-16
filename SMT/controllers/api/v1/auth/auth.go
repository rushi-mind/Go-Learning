package authV1Controllers

import (
	"SMT/config"
	"SMT/models"
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

func AuthAdmin(c *gin.Context) {
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

func AuthTeacher(c *gin.Context) {
	var requestBody requestTypes.AuthTeacher
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utility.ErrorResponse(c, utility.GetErrorMessage(err))
		return
	}
	if !repository.IsValidEmployeeID(requestBody.EmployeeID) {
		utility.ErrorResponse(c, stringTypes.INVALID_TEACHER_ID)
		return
	}
	encryptedPassword := repository.GetTeacherPasswordFromDB(requestBody.EmployeeID)
	if !utility.ValidatePassword(requestBody.Password, encryptedPassword) {
		utility.ErrorResponse(c, stringTypes.INVALID_PASSWORD)
		return
	}
	teacher := repository.GetTeacherByEmployeeID(requestBody.EmployeeID)
	token, err := utility.CreateJWT(map[string]interface{}{
		"id":            teacher.ID,
		"employee_id":   teacher.EmployeeID,
		"email":         teacher.Email,
		"department_id": teacher.DepartmentID,
		"role":          "teacher",
	})
	if err != nil {
		utility.ErrorResponse(c, stringTypes.LOGIN_FAILED)
		return
	}

	utility.SuccessResponseWithToken(c, token)

}
