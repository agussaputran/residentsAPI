package main

import (
	"fmt"
	"testcasetwo/config"
	"testcasetwo/controllers"
	"testcasetwo/middleware"
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
	models.SeederUser(pgDB)

	authMiddleware := middleware.Auth

	router := gin.Default()

	router.POST("/auth/login", strDB.LoginUser)

	router.POST("/province", authMiddleware, strDB.PostCreateProvince)
	router.GET("/province", authMiddleware, strDB.GetReadProvince)
	router.PATCH("/province", authMiddleware, strDB.PatchUpdateProvince)
	router.DELETE("/province", authMiddleware, strDB.DeleteRemoveProvince)
	router.POST("/upload/:id", authMiddleware, strDB.UploadSingle)

	router.POST("/district", authMiddleware, strDB.PostCreateDistrict)
	router.GET("/district", authMiddleware, strDB.GetReadDistrict)
	router.PATCH("/district", authMiddleware, strDB.PatchUpdateDistrict)
	router.DELETE("/district", authMiddleware, strDB.DeleteRemoveDistrict)

	router.POST("/subdistrict", authMiddleware, strDB.PostCreateSubDistrict)
	router.GET("/subdistrict", authMiddleware, strDB.GetReadSubDistrict)
	router.PATCH("/subdistrict", authMiddleware, strDB.PatchUpdateSubDistrict)
	router.DELETE("/subdistrict", authMiddleware, strDB.DeleteRemoveSubDistrict)

	router.POST("/person", authMiddleware, strDB.PostCreatePerson)
	router.GET("/person", authMiddleware, strDB.GetReadPerson)
	router.PATCH("/person", authMiddleware, strDB.PatchUpdatePerson)
	router.DELETE("/person", authMiddleware, strDB.DeleteRemovePerson)

	router.Run()
}
