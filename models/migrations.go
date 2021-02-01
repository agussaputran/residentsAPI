package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var (
		checkProvinces    bool
		checkDistricts    bool
		checkSubDistricts bool
		checkPersons      bool
	)

	db.Migrator().DropTable(&Provinces{})
	db.Migrator().DropTable(&Districts{})
	db.Migrator().DropTable(&SubDistricts{})
	db.Migrator().DropTable(&Persons{})

	checkProvinces = db.Migrator().HasTable(&Provinces{})
	if !checkProvinces {
		db.Migrator().CreateTable(&Provinces{})
		fmt.Println("Create Provinces Table")
	}

	checkDistricts = db.Migrator().HasTable(&Districts{})
	if !checkDistricts {
		db.Migrator().CreateTable(&Districts{})
		fmt.Println("Create Districts Table")
	}

	checkSubDistricts = db.Migrator().HasTable(&SubDistricts{})
	if !checkSubDistricts {
		db.Migrator().CreateTable(&SubDistricts{})
		fmt.Println("Create SubDistricts Table")
	}

	checkPersons = db.Migrator().HasTable(&Persons{})
	if !checkPersons {
		db.Migrator().CreateTable(&Persons{})
		fmt.Println("Create Persons Table")
	}
}
