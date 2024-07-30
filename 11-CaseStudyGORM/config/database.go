package config

import (
	"CaseStudyGorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConfig() *gorm.DB {
	// Koneksi ke database MySQL
	dsn := "root:root@tcp(127.0.0.1:3306)/TestGorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrasi tabel
	db.AutoMigrate(&model.Kelas{}, &model.Siswa{}, &model.Pelajaran{}, &model.Nilai{})
	return db
}
