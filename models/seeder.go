package models

import (
	"fmt"

	"gorm.io/gorm"
)

// SeederProvince func
func SeederProvince(db *gorm.DB) {
	var provinceArr = [...]string{
		"Jawa Tengah",
		"DI Yogyakarta",
		"Jawa Barat",
		"Jawa Timur",
	}

	var prov Provinces

	for _, v := range provinceArr {
		prov.Name = v
		prov.ID = 0
		db.Create(&prov)
	}
	fmt.Println("Seeder Prov created")
}

// SeederDistrict func
// func SeederDistrict(db *gorm.DB) {
// 	var districtArr = [...][2]string{
// 		{"1", "Kota Jogja"},
// 		{"1", "Bantul"},
// 		{"1", "Sleman"},
// 		{"1", "Kulon Progo"},
// 		{"1", "Gunung Kidul"},
// 	}

// 	var dist Districts
// 	for _, v := range districtArr{
// 		dist.ProvinceID = v
// 	}
// }
