package controller

import (
	"CaseStudyGorm/dto"
	"CaseStudyGorm/model"
	"fmt"
)

func InsertKelas() {
	kelas := model.Kelas{
		ID:   1,
		Nama: "X-A",
	}

	dto.InsertKelas(kelas)
}

func InsertSiswa() {
	siswa := model.Siswa{
		ID:      1,
		Nama:    "Budi",
		Umur:    15,
		KelasID: 1,
	}

	siswa1 := model.Siswa{
		ID:      2,
		Nama:    "Adi",
		Umur:    15,
		KelasID: 1,
	}

	dto.InsertSiswa(siswa)
	dto.InsertSiswa(siswa1)
}

func InsertPelajaran() {
	pelajaran := model.Pelajaran{
		ID:   1,
		Nama: "Matematika",
	}

	pelajaran1 := model.Pelajaran{
		ID:   2,
		Nama: "Bahasa",
	}

	dto.InsertPelajaran(pelajaran)
	dto.InsertPelajaran(pelajaran1)
}

func InsertNilai() {
	nilai := model.Nilai{
		ID:          3,
		SiswaID:     2,
		PelajaranID: 1,
		Nilai:       100,
	}

	nilai1 := model.Nilai{
		ID:          4,
		SiswaID:     1,
		PelajaranID: 2,
		Nilai:       90,
	}

	dto.InsertNilai(nilai)
	dto.InsertNilai(nilai1)
}

func ShowSiswa() {
	data := dto.GetAllSiswa()
	for _, val := range data {
		fmt.Println(val.ID, "", val.Nama)
		fmt.Println("Kelas ", val.Kelas.ID, "", val.Kelas.Nama)
	}
}

func ShowKelas() {
	data := dto.GetAllKelas()
	for _, val := range data {
		fmt.Println(val.ID, "", val.Nama)
		for _, valIn := range val.Siswa {
			fmt.Println(valIn.ID, "", valIn.Nama, "", valIn.Umur)
		}
	}
}

func ShowNilaiBySiswa() {
	data := dto.GetNilaiBySiswa()
	for _, val := range data {
		fmt.Println(val.Nama, "", val.Umur)
		fmt.Println("Nilai")
		for _, valIn := range val.Nilai {
			fmt.Println("-", dto.GetPelajaranByID(int(valIn.PelajaranID)).Nama, "", valIn.Nilai)
		}
	}
}

func ShowAverageNilaiSiswa() {
	data := dto.GetNilaiBySiswa()
	for _, val := range data {
		fmt.Println(val.Nama, "", val.Umur)
		fmt.Println("Nilai : ", dto.GetAverageBySiswa(int(val.ID)))
	}
}
