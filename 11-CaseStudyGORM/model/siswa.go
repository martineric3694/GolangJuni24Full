package model

type Siswa struct {
	ID        uint `gorm:"primaryKey"`
	Nama      string
	Umur      uint
	KelasID   uint        `gorm:"notNull"`
	Kelas     Kelas       `gorm:"foreignKey:KelasID"`
	Nilai     []Nilai     `gorm:"foreignKey:SiswaID"`
	Pelajaran []Pelajaran `gorm:"many2many:nilais;"`
}
