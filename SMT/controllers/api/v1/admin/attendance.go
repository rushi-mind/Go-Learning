package adminV1Controller

import (
	"SMT/models"
	"SMT/repository"
	stringTypes "SMT/types/strings"
	"SMT/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func AddAttendance(c *gin.Context) {
	var requestBody []models.Attendace
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Default().Println(err)
		utility.ErrorResponse(c, err.Error())
		return
	}
	if err = repository.InsertAttendance(requestBody); err != nil {
		log.Default().Println(err)
		utility.ErrorResponse(c, stringTypes.FAILED_INSERT_ATTENDANCE)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.ATTENDANCE_INSERTED)
}
