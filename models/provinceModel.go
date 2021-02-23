package models

import (
	"time"

	"gorm.io/gorm"
)

// Provinces model
type Provinces struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `json:"province_name"`
	District  []Districts    `gorm:"ForeignKey:ProvinceID" json:"districts"`
}
