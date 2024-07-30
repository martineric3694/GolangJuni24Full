package dto

import (
	"CaseStudyGorm/config"
	"CaseStudyGorm/model"
	"fmt"
)

func InsertNilai(nilai model.Nilai) {
	db := config.DBConfig()
	db.Create(nilai)
	fmt.Println("Insert Successfully")
}

func GetAllNilai() []model.Nilai {
	db := config.DBConfig()
	var result []model.Nilai
	db.Model(&model.Nilai{}).Preload("Siswa").Preload("Pelajaran").Find(&result)
	return result
}

func GetAverageBySiswa(id int) float64 {
	db := config.DBConfig()
	var result float64
	db.Model(&model.Nilai{}).Where("siswa_id=?", id).Select("AVG(nilai)").First(&result)
	return result
}
