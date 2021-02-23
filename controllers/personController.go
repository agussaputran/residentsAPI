package controllers

import (
	"fmt"
	"net/http"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
)

type personResponse struct {
	ID           uint   `json:"id"`
	FullName     string `json:"full_name"`
	FirstName    string `json:"firts_name"`
	LastName     string `json:"last_name"`
	BirthDate    string `json:"birth_date"`
	BirthPlace   string `json:"birth_place"`
	Gender       string `json:"gender"`
	ZoneLocation string `json:"zone_location"`
	Subdistrict  string `json:"sub_district_name"`
	District     string `json:"district_name"`
	Province     string `json:"province_name"`
	Photo        string `json:"photo_url"`
}

// PostCreatePerson route struct method
func (strDB *StrDB) PostCreatePerson(c *gin.Context) {
	var (
		person models.Persons
		result gin.H
	)

	if err := c.Bind(&person); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Create(&person)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"id":              person.ID,
				"nip":             person.Nip,
				"full_name":       person.FullName,
				"firts_name":      person.FirstName,
				"last_name":       person.LastName,
				"sub_district_id": person.SubDistrictID,
				"birth_date":      person.BirthDate,
				"birth_place":     person.BirthPlace,
				"gender":          person.Gender,
				"zone_location":   person.ZoneLocation,
				"photo_url":       person.PhotoUrl,
				"created_at":      person.CreatedAt,
				"update_at":       person.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadPerson route func
func (strDB *StrDB) GetReadPerson(c *gin.Context) {
	var (
		person   []models.Persons
		response []personResponse
		result   gin.H
	)

	strDB.DB.Model(&person).Select(`persons.id, persons.full_name, persons.first_name,
	persons.last_name, persons.birth_date, persons.birth_place,
	persons.gender, persons.zone_location, persons.photo_url as photo, sub_districts.name as subdistrict,
	districts.name as district, provinces.name as province`).Joins(`left join sub_districts
	on sub_districts.id = persons.sub_district_id left join districts on districts.id =
	sub_districts.district_id left join provinces on provinces.id = districts.province_id`).Scan(&response)
	if length := len(response); length <= 0 {
		result = ResultAPINilResponse(response, length)
	} else {
		result = ResultAPIResponse(response, length)
	}

	c.JSON(http.StatusOK, result)
}

// PatchUpdatePerson route struct method
func (strDB *StrDB) PatchUpdatePerson(c *gin.Context) {
	var (
		person models.Persons
		result gin.H
	)

	id := c.Query("id")

	if err := c.Bind(&person); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		strDB.DB.Model(&person).Updates(models.Persons{FullName: person.FullName, FirstName: person.FirstName, LastName: person.LastName, SubDistrictID: person.SubDistrictID, ZoneLocation: person.ZoneLocation}).Where("id = ?", id)
		result = gin.H{
			"message": "success",
			"data": map[string]interface{}{
				"fullName":      person.FullName,
				"firstName":     person.FirstName,
				"lastName":      person.LastName,
				"subDistrictID": person.SubDistrictID,
				"birthDate":     person.BirthDate,
				"birthPlace":    person.BirthPlace,
				"gender":        person.Gender,
				"zoneLocation":  person.ZoneLocation,
				"photo_url":     person.PhotoUrl,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// DeleteRemovePerson route struct method
func (strDB *StrDB) DeleteRemovePerson(c *gin.Context) {
	var (
		person models.Persons
		result gin.H
	)

	id := c.Query("id")
	strDB.DB.Delete(&person, id)
	result = gin.H{
		"Message": "Success delete district",
	}
	c.JSON(http.StatusOK, result)
}
