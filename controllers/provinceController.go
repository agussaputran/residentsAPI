package controllers

import (
	"fmt"
	"net/http"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
)

type provinceResponse struct {
	ID       uint
	Province string
}

// PostCreateProvince route struct method
func (strDB *StrDB) PostCreateProvince(c *gin.Context) {
	var (
		province models.Provinces
		result   gin.H
	)

	if err := c.Bind(&province); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&province)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":         province.ID,
				"province":   province.Name,
				"created_at": province.CreatedAt,
				"update_at":  province.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadProvince func route
func (strDB *StrDB) GetReadProvince(c *gin.Context) {
	var (
		province []models.Provinces
		response []provinceResponse
		result   gin.H
	)

	strDB.DB.Model(&province).Select("id, name as province").Scan(&response)
	if length := len(response); length <= 0 {
		result = ResultAPINilResponse(response, length)
	} else {
		result = ResultAPIResponse(response, length)
	}

	c.JSON(http.StatusOK, result)
}

// PatchUpdateProvince route struct method
func (strDB *StrDB) PatchUpdateProvince(c *gin.Context) {
	var (
		province models.Provinces
		result   gin.H
	)

	id := c.Query("id")

	if err := c.Bind(&province); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Model(&province).Where("id = ?", id).Update("name", province.Name)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"province": province.Name,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// DeleteRemoveProvince route struct method
func (strDB *StrDB) DeleteRemoveProvince(c *gin.Context) {
	var (
		province models.Provinces
		result   gin.H
	)

	id := c.Query("id")
	strDB.DB.Delete(&province, id)
	result = gin.H{
		"Message": "Success delete",
	}
	c.JSON(http.StatusOK, result)
}
