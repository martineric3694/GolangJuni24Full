package model

type Kelas struct {
	ID    uint `gorm:"primaryKey"`
	Nama  string
	Siswa []Siswa `gorm:"foreignKey:KelasID"`
}
