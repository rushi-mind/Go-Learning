package admin

import (
	"SMT/types/responses"
	"SMT/types/strings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(c *gin.Context, message string) {
	response := responses.AuthError{
		Base: responses.Base{
			Status:  0,
			Message: message,
		},
	}
	c.JSON(http.StatusBadRequest, response)
}

func SendLoginSuccessResponse(c *gin.Context, token string) {
	c.JSON(http.StatusOK, responses.AdminAuthSuccess{
		Base: responses.Base{
			Status:  1,
			Message: strings.SUCCESS_LOGIN,
		},
		Data: responses.TokenResponse{
			Token: token,
		},
	})
}

func SendSuccessResponseWithoutBody(c *gin.Context, message string) {
	c.JSON(http.StatusOK, responses.SuccessResponse{
		Base: responses.Base{
			Status:  1,
			Message: message,
		},
	})
}
