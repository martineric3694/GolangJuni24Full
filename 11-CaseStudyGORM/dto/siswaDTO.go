package dto

import (
	"CaseStudyGorm/config"
	"CaseStudyGorm/model"
	"fmt"
)

func InsertSiswa(siswa model.Siswa) {
	db := config.DBConfig()
	db.Create(&siswa)
	fmt.Println("Insert Successfully")
}

func GetAllSiswa() []model.Siswa {
	db := config.DBConfig()
	var result []model.Siswa
	db.Model(&model.Siswa{}).Preload("Kelas").Find(&result)
	return result
}

func GetNilaiBySiswa() []model.Siswa {
	db := config.DBConfig()
	var result []model.Siswa
	db.Model(&model.Siswa{}).Preload("Nilai").Preload("Pelajaran").Find(&result)
	return result
}
