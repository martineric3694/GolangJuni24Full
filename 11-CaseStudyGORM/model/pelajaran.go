package model

type Pelajaran struct {
	ID    uint `gorm:"primaryKey"`
	Nama  string
	Siswa []Siswa `gorm:"many2many:nilais;"`
}
