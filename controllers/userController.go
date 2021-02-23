package controllers

import (
	"log"
	"net/http"
	"os"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser func
func (strDB *StrDB) LoginUser(c *gin.Context) {
	var (
		user   models.Users
		userDB models.Users
		result gin.H
	)

	guestToken := os.Getenv("TOKEN_GUEST")
	adminToken := os.Getenv("TOKEN_ADMIN")

	if err := c.Bind(&user); err != nil {
		log.Println("Data tidak ada, error : ", err.Error())
	}

	strDB.DB.Where("email = ?", user.Email).First(&userDB)

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		log.Println("Email ", user.Email, " Password salah")
		result = gin.H{
			"message": "email atau password salah",
		}
	} else {
		if userDB.Role == "admin" {
			result = gin.H{
				"message": "anda berhasil login",
				"token":   adminToken,
			}
		} else {
			result = gin.H{
				"message": "anda berhasil login",
				"token":   guestToken,
			}
		}
	}

	c.JSON(http.StatusOK, result)
}
