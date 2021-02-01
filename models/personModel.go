package models

import (
	"gorm.io/gorm"
)

// Persons model
type Persons struct {
	gorm.Model
	Nip, FullName, FirstName, LastName, BirthDate, BirthPlace, Gender, ZoneLocation, PhotoUrl string
	SubDistrictID                                                                             uint
}
