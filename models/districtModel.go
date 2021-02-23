package models

import (
	"time"

	"gorm.io/gorm"
)

// Districts model
type Districts struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `json:"district_name"`
	ProvinceID  uint           `json:"provice_id"`
	SubDistrict []SubDistricts `gorm:"ForeignKey:DistrictID" json:"sub_districts"`
}
