package main

import (
	"testcasetwo/config"
	"testcasetwo/controllers"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// db connect to postgre DB
	pgDB := config.Connect()
	strDB := controllers.StrDB{DB: pgDB}

	// Migrations
	models.Migrations(pgDB)

	router := gin.Default()

	router.POST("/province", strDB.PostCreateProvince)
	router.GET("/province", strDB.GetReadProvince)
	router.PATCH("/province", strDB.PatchUpdateProvince)
	router.DELETE("/province", strDB.DeleteRemoveProvince)
	router.POST("/upload", controllers.UploadSingle)

	router.POST("/district", strDB.PostCreateDistrict)
	router.GET("/district", strDB.GetReadDistrict)
	router.PATCH("/district", strDB.PatchUpdateDistrict)
	router.DELETE("/district", strDB.DeleteRemoveDistrict)

	router.POST("/subdistrict", strDB.PostCreateSubDistrict)
	router.GET("/subdistrict", strDB.GetReadSubDistrict)
	router.PATCH("/subdistrict", strDB.PatchUpdateSubDistrict)
	router.DELETE("/subdistrict", strDB.DeleteRemoveSubDistrict)

	router.POST("/person", strDB.PostCreatePerson)
	router.GET("/person", strDB.GetReadPerson)
	router.PATCH("/person", strDB.PatchUpdatePerson)
	router.DELETE("/person", strDB.DeleteRemovePerson)

	router.Run()
}
