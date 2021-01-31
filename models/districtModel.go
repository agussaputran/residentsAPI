package models

import "gorm.io/gorm"

// Districts model
type Districts struct {
	gorm.Model
	Name        string
	ProvinceID  uint
	Province    Provinces
	SubDistrict []SubDistricts `gorm:"ForeignKey:DistrictID"`
}
