package main

import (
	"fmt"
	"testcasetwo/config"
	"testcasetwo/controllers"
	"testcasetwo/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found ..")
	}
}

func main() {
	// db connect to postgre DB
	pgDB := config.Connect()
	strDB := controllers.StrDB{DB: pgDB}

	// Migrations
	models.Migrations(pgDB)

	// Seed to postgre DB
	models.SeederProvince(pgDB)
	models.SeederDistrict(pgDB)
	models.SeederSubDistrict(pgDB)

	router := gin.Default()

	router.POST("/province", strDB.PostCreateProvince)
	router.GET("/province", strDB.GetReadProvince)
	router.PATCH("/province", strDB.PatchUpdateProvince)
	router.DELETE("/province", strDB.DeleteRemoveProvince)
	router.POST("/upload/:id", strDB.UploadSingle)

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
