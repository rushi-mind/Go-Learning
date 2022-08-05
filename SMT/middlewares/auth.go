package middlewares

import (
	stringTypes "SMT/types/strings"
	"SMT/utility"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || len(token) == 0 {
		c.String(401, "token not found")
		c.Abort()
		return
	}
	payload, err := utility.GetPayloadFromToken(strings.Split(token, " ")[1])
	if err != nil || payload["adminId"] == "" {
		utility.ErrorResponse(c, stringTypes.AUTHENTICATION_FAILED)

		return
	}
	c.Set("user", payload)
	c.Next()
}
