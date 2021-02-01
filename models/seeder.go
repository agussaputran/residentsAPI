package models

import (
	"fmt"
	"strconv"

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
func SeederDistrict(db *gorm.DB) {
	var districtArr = [...][2]string{
		{"2", "Kota Jogja"},
		{"2", "Bantul"},
		{"2", "Sleman"},
		{"1", "Klaten"},
		{"1", "Magelang"},
	}

	var dist Districts
	for _, v := range districtArr {
		data, _ := strconv.ParseUint(v[0], 10, 64)
		dist.ProvinceID = uint(data)
		dist.Name = v[1]
		dist.ID = 0
		db.Create(&dist)
	}
	fmt.Println("Seeder District created")
}

//SeederSubDistrict func
func SeederSubDistrict(db *gorm.DB) {
	var subDistrictArray = [...][2]string{
		{"1", "Umbulharjo"},
		{"2", "Banguntapan"},
		{"3", "Ngaglik"},
		{"4", "Klaten Selatan"},
		{"5", "Muntilan"},
	}

	var subDist SubDistricts
	for _, v1 := range subDistrictArray {
		data, _ := strconv.ParseUint(v1[0], 10, 64)
		subDist.DistrictID = uint(data)
		subDist.Name = v1[1]
		subDist.ID = 0
		db.Create(&subDist)

	}
	fmt.Println("Seeder Sub District created")
}
