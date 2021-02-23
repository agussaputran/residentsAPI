package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthUser middleware with static token
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

// Auth func with jwt
func Auth(c *gin.Context) {
	secret := os.Getenv("SECRET_TOKEN")
	tokenStringHeader := c.Request.Header.Get("Authorization")
	allowedMethod := c.Request.Method
	token, err := jwt.Parse(tokenStringHeader, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Method tidak diketahui atau bukan HS256 %V", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if token != nil && err == nil {
		payload := token.Claims.(jwt.MapClaims)
		log.Println("Token Verified")

		if payload["role"] != "admin" && allowedMethod != "GET" {
			result := gin.H{
				"code":    401,
				"status":  "Unauthorized",
				"message": "Tidak punya akses",
				"data":    nil,
			}
			LogTerminal(c)
			// LogSentry(c)
			c.Abort()
			c.JSON(http.StatusUnauthorized, result)
		} else {
			// LogUser(payload)
			LogTerminal(c)
		}

	} else if err != nil {
		log.Println("Wrong Token, error -> ", err.Error())
		result := gin.H{
			"code":    401,
			"status":  "Unauthorized",
			"message": "Token tidak valid",
			"data":    nil,
		}
		LogTerminal(c)
		c.Abort()
		c.JSON(http.StatusUnauthorized, result)
	}
}
