package models

import (
	"time"

	"gorm.io/gorm"
)

// SubDistricts model
type SubDistricts struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name       string         `json:"sub_district_name"`
	DistrictID uint           `json:"district_id"`
	Person     []Persons      `gorm:"ForeignKey:SubDistrictID" json:"persons"`
}
