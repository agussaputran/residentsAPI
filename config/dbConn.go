package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnvVar(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// Connect with gorm
func Connect() *gorm.DB {
	var userDB, pwDB, portDB, hostDB, nameDB string
	userDB = getEnvVar("DB_USER")
	pwDB = getEnvVar("DB_PASSWORD")
	portDB = getEnvVar("DB_PORT")
	hostDB = getEnvVar("DB_HOST")
	nameDB = getEnvVar("DB_NAME")

	conn := " host=" + hostDB +
		" user=" + userDB +
		" password=" + pwDB +
		" dbname=" + nameDB +
		" port=" + portDB +
		" sslmode=disable TimeZone=Asia/Shanghai"

	db, errConn := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if errConn != nil {
		panic("failed to connect to the database")
	} else {
		fmt.Println("successful connection")
	}
	return db
}
