package dto

import (
	"CaseStudyGorm/config"
	"CaseStudyGorm/model"
	"fmt"
)

func InsertKelas(kelas model.Kelas) {
	db := config.DBConfig()
	db.Create(kelas)
	fmt.Println("Insert Successfully")
}

func GetAllKelas() []model.Kelas {
	db := config.DBConfig()
	var result []model.Kelas
	db.Model(&model.Kelas{}).Preload("Siswa").Find(&result)
	return result
}
