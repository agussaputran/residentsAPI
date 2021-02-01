package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
)

// UploadSingle func
func (strDB *StrDB) UploadSingle(c *gin.Context) {
	var (
		person models.Persons
		result gin.H
	)

	file, _ := c.FormFile("photo")
	fileName := strings.ToLower(file.Filename)
	photoURL := "http://localhost:8080" + strings.ReplaceAll(fileName, " ", "%")

	id := c.Param("id")
	if err := c.Bind(&person); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Model(&person).Where("id = ?", id).Update("photo_url", photoURL)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"photo_url": person.PhotoUrl,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}
