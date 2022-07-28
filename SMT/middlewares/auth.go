package middlewares

import (
	adminServices "SMT/services/admin"
	messageStrings "SMT/types/strings"
	"SMT/utility"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.String(401, "token not entered")
		c.Abort()
	}
	payload, err := utility.GetPayloadFromToken(strings.Split(token, " ")[1])
	if err != nil || payload["adminId"] == "" {
		adminServices.SendErrorResponse(c, messageStrings.AUTHENTICATION_FAILED)
	}
	c.Set("user", payload)
	c.Next()
}
