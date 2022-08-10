package middlewares

import (
	stringTypes "SMT/types/strings"
	"SMT/utility"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddlewareAdmin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || len(token) == 0 {
		c.String(401, "token not found")
		c.Abort()
		return
	}
	payload, err := utility.GetPayloadFromToken(strings.Split(token, " ")[1])
	if err != nil || payload["role"] != "admin" {
		utility.ErrorResponse(c, stringTypes.AUTHORIZATION_FAILED)
		c.Abort()
		return
	}
	c.Set("user", payload)
	c.Next()
}

func AuthMiddlewareStudent(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || len(token) == 0 {
		c.String(401, "token not found")
		c.Abort()
		return
	}
	payload, err := utility.GetPayloadFromToken(strings.Split(token, " ")[1])
	if err != nil || payload["role"] != "student" {
		utility.ErrorResponse(c, stringTypes.AUTHORIZATION_FAILED)
		c.Abort()
		return
	}
	c.Set("user", payload)
	c.Next()
}
