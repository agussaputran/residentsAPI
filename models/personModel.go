package models

import (
	"time"

	"gorm.io/gorm"
)

// Persons model
type Persons struct {
	gorm.Model
	FullName, FirstName, LastName, BirthPlace, Gender, ZoneLocation string
	BirthDate                                                       time.Time
	SubDistrictID                                                   uint
}
