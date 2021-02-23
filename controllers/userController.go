package controllers

import (
	"log"
	"net/http"
	"os"
	"testcasetwo/models"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	if err := c.Bind(&user); err != nil {
		log.Println("Data tidak ada, error : ", err.Error())
	}

	strDB.DB.Where("email = ?", user.Email).First(&userDB)

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		log.Println("Email ", user.Email, " Password salah")
		result = gin.H{
			"code":    401,
			"status":  "Unauthorized",
			"message": "Email atau password salah",
			"data":    nil,
		}
	} else {
		type authCustomClaims struct {
			ID    uint   `json:"id"`
			Email string `json:"email"`
			Role  string `json:"role"`
			jwt.StandardClaims
		}

		claims := &authCustomClaims{
			userDB.ID,
			userDB.Email,
			userDB.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}
		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, err := sign.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
		if err != nil {
			log.Println("Gagal create token, message ", err.Error())
			result = gin.H{
				"code":    400,
				"status":  "Bad Request",
				"message": "Anda gagal login",
				"data": map[string]interface{}{
					"token": nil,
				},
			}
		} else {
			log.Println("Email ", user.Email, " Berhasil login")
			result = gin.H{
				"code":    200,
				"status":  "Ok",
				"message": "Anda berhasil login",
				"data": map[string]interface{}{
					"token": token,
				},
			}
		}
	}

	c.JSON(http.StatusOK, result)
}
