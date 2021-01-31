package models

import (
	"gorm.io/gorm"
)

// Persons model
type Persons struct {
	gorm.Model
	Nip, FullName, FirstName, LastName, BirthPlace, Gender, ZoneLocation string
	BirthDate                                                            string
	SubDistrictID                                                        uint
	SubDistrict                                                          SubDistricts
}
