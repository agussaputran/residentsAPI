package controllers

import (
	"fmt"
	"net/http"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
)

// PostCreateDistrict route struct method
func (strDB *StrDB) PostCreateDistrict(c *gin.Context) {
	var (
		district models.Districts
		result   gin.H
	)

	if err := c.Bind(&district); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&district)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":          district.ID,
				"province_id": district.ProvinceID,
				"district":    district.Name,
				"created_at":  district.CreatedAt,
				"update_at":   district.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadDistrict route func
func (strDB *StrDB) GetReadDistrict(c *gin.Context) {
	var (
		district []models.Districts
		result   gin.H
	)

	strDB.DB.Find(&district)
	if length := len(district); length <= 0 {
		result = ResultAPINilResponse(district, length)
	} else {
		result = ResultAPIResponse(district, length)
	}

	c.JSON(http.StatusOK, result)
}

// PatchUpdateDistrict route struct method
func (strDB *StrDB) PatchUpdateDistrict(c *gin.Context) {
	var (
		district models.Districts
		result   gin.H
	)

	id := c.Query("id")

	if err := c.Bind(&district); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Model(&district).Where("id = ?", id).Update("name", district.Name)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"province": district.Name,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// DeleteRemoveDistrict route struct method
func (strDB *StrDB) DeleteRemoveDistrict(c *gin.Context) {
	var (
		district models.Districts
		result   gin.H
	)

	id := c.Query("id")
	strDB.DB.Delete(&district, id)
	result = gin.H{
		"Message": "Success delete district",
	}
	c.JSON(http.StatusOK, result)
}
