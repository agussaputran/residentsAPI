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

		constraintDistricts    bool
		constraintSubDistricts bool
		constraintPersons      bool
	)

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

	// check contraint province to district table
	constraintDistricts = db.Migrator().HasConstraint(&Provinces{}, "Districts")
	if !constraintDistricts {
		db.Migrator().CreateConstraint(&Provinces{}, "Distritcs")
		fmt.Println("Create Constraint Provinces|Districts")
	}

	// check contraint district to subDistrict table
	constraintSubDistricts = db.Migrator().HasConstraint(&Districts{}, "SubDistricts")
	if !constraintSubDistricts {
		db.Migrator().CreateConstraint(&Districts{}, "SubDistricts")
		fmt.Println("Create Constraint Districts|SubDistricts")
	}

	// check contraint subDistrict to Person table
	constraintPersons = db.Migrator().HasConstraint(&SubDistricts{}, "Persons")
	if !constraintPersons {
		db.Migrator().CreateConstraint(&SubDistricts{}, "Persons")
		fmt.Println("Create Constraint SubDistricts|Persons")
	}
}
