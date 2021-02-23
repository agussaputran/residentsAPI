package models

import (
	"gorm.io/gorm"
)

// Persons model
type Persons struct {
	gorm.Model
	Nip           string `json:"nip"`
	FullName      string `json:"full_ name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	BirthDate     string `json:"birth_Date"`
	BirthPlace    string `json:"birth_place"`
	Gender        string `json:"gender"`
	ZoneLocation  string `json:"zonce_location"`
	PhotoURL      string `json:"photo_url"`
	SubDistrictID uint   `json:"sub_district_id"`
}
