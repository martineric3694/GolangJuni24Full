package dto

import (
	"CaseStudyGorm/config"
	"CaseStudyGorm/model"
	"fmt"
)

func InsertPelajaran(pelajaran model.Pelajaran) {
	db := config.DBConfig()
	db.Create(pelajaran)
	fmt.Println("Insert Successfully")
}

func GetAllPelajaran() []model.Pelajaran {
	db := config.DBConfig()
	var result []model.Pelajaran
	db.Model(&model.Pelajaran{}).Preload("Siswa").Find(&result)
	return result
}

func GetPelajaranByID(id int) model.Pelajaran {
	db := config.DBConfig()
	var result model.Pelajaran
	db.Model(&model.Pelajaran{}).Where(id).First(&result)
	return result
}
