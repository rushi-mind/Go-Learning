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

func SuccessResponseWithToken(c *gin.Context, token string) {
	c.JSON(http.StatusOK, responseTypes.TokenResponse{
		Token: token,
	})
}

func SuccessResponseWithData(c *gin.Context, message string, data any, count int) {
	c.JSON(http.StatusOK, responseTypes.SuccessResponseData{
		Meta: responseTypes.Meta{Base: responseTypes.Base{Status: 1, Message: message}, Count: count},
		Data: data,
	})
}

func SuccessResponseWithoutData(c *gin.Context, message string) {
	c.JSON(http.StatusOK, responseTypes.Base{
		Status:  1,
		Message: message,
	})
}
