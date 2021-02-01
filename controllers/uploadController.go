package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadSingle func
func UploadSingle(c *gin.Context) {
	file, _ := c.FormFile("photo-profile")

	// fmt.Println(file.Filename)
	// fmt.Println(file.Size)
	// fmt.Println(file.Header)

	c.JSON(http.StatusOK, gin.H{
		"file_name":   file.Filename,
		"file_size":   file.Size,
		"file_header": file.Header,
	})
}
