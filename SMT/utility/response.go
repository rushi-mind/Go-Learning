package utility

import (
	responseTypes "SMT/types/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, responseTypes.Base{
		Status:  0,
		Message: message,
	})
}

func SuccessResponseWithData(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, responseTypes.SuccessResponseData{
		Base: responseTypes.Base{Status: 1, Message: message},
		Data: data,
	})
}

func SuccessResponseWithoutData(c *gin.Context, message string) {
	c.JSON(http.StatusOK, responseTypes.Base{
		Status:  1,
		Message: message,
	})
}
