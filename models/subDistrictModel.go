package models

import "gorm.io/gorm"

// SubDistricts model
type SubDistricts struct {
	gorm.Model
	Name       string
	DistrictID uint
	District   Districts
	Person     []Persons `gorm:"ForeignKey:SubDistrictID"`
}
