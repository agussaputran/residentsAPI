package controllers

import (
	"fmt"
	"net/http"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
)

type subDistrictResponse struct {
	ID          uint
	Subdistrict string
	District    string
	Province    string
}

// PostCreateSubDistrict route struct method
func (strDB *StrDB) PostCreateSubDistrict(c *gin.Context) {
	var (
		subDistrict models.SubDistricts
		result      gin.H
	)

	if err := c.Bind(&subDistrict); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&subDistrict)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"ID":          subDistrict.ID,
				"province_id": subDistrict.DistrictID,
				"subDistrict": subDistrict.Name,
				"created_at":  subDistrict.CreatedAt,
				"update_at":   subDistrict.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadSubDistrict route func
func (strDB *StrDB) GetReadSubDistrict(c *gin.Context) {
	var (
		subDistrict []models.SubDistricts
		response    []subDistrictResponse
		result      gin.H
	)

	strDB.DB.Model(&subDistrict).Select(`sub_districts.id,
		sub_districts.name as subdistrict,
		districts.name as district,
		provinces.name as province`).Joins(`left join districts
		on districts.id = sub_districts.district_id left join provinces
		on provinces.id = districts.province_id`).Scan(&response)
	if length := len(response); length <= 0 {
		result = ResultAPINilResponse(response, length)
	} else {
		result = ResultAPIResponse(response, length)
	}

	c.JSON(http.StatusOK, result)
}

// PatchUpdateSubDistrict route struct method
func (strDB *StrDB) PatchUpdateSubDistrict(c *gin.Context) {
	var (
		subDistrict models.SubDistricts
		result      gin.H
	)

	id := c.Query("id")

	if err := c.Bind(&subDistrict); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Model(&subDistrict).Where("id = ?", id).Update("name", subDistrict.Name)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"province": subDistrict.Name,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// DeleteRemoveSubDistrict route struct method
func (strDB *StrDB) DeleteRemoveSubDistrict(c *gin.Context) {
	var (
		subDistrict models.SubDistricts
		result      gin.H
	)

	id := c.Query("id")
	strDB.DB.Delete(&subDistrict, id)
	result = gin.H{
		"Message": "Success delete district",
	}
	c.JSON(http.StatusOK, result)
}
