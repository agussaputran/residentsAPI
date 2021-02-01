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
	fileNameLower := strings.ToLower(file.Filename)
	fileName := strings.ReplaceAll(fileNameLower, " ", "%")
	path := "images/" + fileName
	photoURL := "http://localhost:8080/" + path + "/" + fileName

	if err := c.SaveUploadedFile(file, path); err != nil {
		fmt.Println("Terjadi Error", err.Error())
	}

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
