package controllers

import (
	"fmt"
	"net/http"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
)

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
				"ID":            person.ID,
				"Nip":           person.Nip,
				"fullName":      person.FullName,
				"firstName":     person.FirstName,
				"lastName":      person.LastName,
				"subDistrictID": person.SubDistrictID,
				"birthDate":     person.BirthDate,
				"birthPlace":    person.BirthPlace,
				"gender":        person.Gender,
				"zoneLocation":  person.ZoneLocation,
				"created_at":    person.CreatedAt,
				"update_at":     person.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetReadPerson route func
func (strDB *StrDB) GetReadPerson(c *gin.Context) {
	var (
		person []models.Persons
		result gin.H
	)

	strDB.DB.Find(&person)
	if length := len(person); length <= 0 {
		result = ResultAPINilResponse(person, length)
	} else {
		result = ResultAPIResponse(person, length)
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
