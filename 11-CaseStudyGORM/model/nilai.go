package model

type Nilai struct {
	ID          uint      `gorm:"primaryKey"`
	SiswaID     uint      `gorm:"notNull"`
	PelajaranID uint      `gorm:"notNull"`
	Nilai       float64   `gorm:"notNull"`
	Siswa       Siswa     `gorm:"foreignKey:SiswaID"`
	Pelajaran   Pelajaran `gorm:"foreignKey:PelajaranID"`
}
