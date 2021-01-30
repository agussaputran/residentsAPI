package controllers

import "gorm.io/gorm"

// StrDB support struct
type StrDB struct {
	DB *gorm.DB
}
