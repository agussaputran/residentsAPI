package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// AuthUser middleware
func AuthUser(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	allowMethod := c.Request.Method
	adminToken := os.Getenv("TOKEN_ADMIN")

	if token[0] != adminToken {
		if allowMethod != "GET" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":  "ERR_UNAUTHORIZED",
				"err_code": http.StatusUnauthorized,
			})
			LogTerminal(c)
			LogSentry(c)
			c.Abort()
		} else {
			LogTerminal(c)
		}
	} else {
		LogTerminal(c)

	}
}
