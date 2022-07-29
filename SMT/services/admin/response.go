package adminServices

import (
	responseTypes "SMT/types/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(c *gin.Context, message string) {
	response := responseTypes.AuthError{
		Base: responseTypes.Base{
			Status:  0,
			Message: message,
		},
	}
	c.JSON(http.StatusBadRequest, response)
}

func SuccessResponseWithData(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, responseTypes.SuccessResponseData{
		Base: responseTypes.Base{Status: 1, Message: message},
		Data: data,
	})
}

func SendLoginSuccessResponse(c *gin.Context, token string) {
	c.JSON(http.StatusOK, responseTypes.AdminAuthSuccess{
		Base: responseTypes.Base{
			Status:  1,
			Message: responseTypes.SUCCESS_LOGIN,
		},
		Data: responseTypes.TokenResponse{
			Token: token,
		},
	})
}

func SendSuccessResponseWithoutBody(c *gin.Context, message string) {
	c.JSON(http.StatusOK, responseTypes.SuccessResponse{
		Base: responseTypes.Base{
			Status:  1,
			Message: message,
		},
	})
}
