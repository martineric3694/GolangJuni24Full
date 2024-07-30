package main

import (
	"database/sql"
	"fmt"
)

type Siswa struct {
	id    int
	nama  string
	kelas string
}

func getAllSiswa(db *sql.DB) {
	// db := connectDB()

	rows, error := db.Query("SELECT * FROM Siswa")
	if error != nil {
		fmt.Println(error.Error())
	}

	var result []Siswa
	var i int = 0
	for rows.Next() {
		i++
		siswa := Siswa{}
		rows.Scan(&siswa.id, &siswa.nama, &siswa.kelas)
		result = append(result, siswa)
	}

	fmt.Println("Total iterasi : ", i)

	for _, value := range result {
		fmt.Println(value.id, " ", value.nama, " ", value.kelas)
	}

	// defer db.Close()
	defer rows.Close()
}

func insertSiswa(db *sql.DB, siswa Siswa) {

	result, error := db.Exec("INSERT INTO Siswa(siswaId,namaSiswa,kelas) VALUES(?,?,?)", siswa.id, siswa.nama, siswa.kelas)
	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println("INSERT SUCCESSFULLY")
	fmt.Print(result.RowsAffected())
	fmt.Println(" affected")
}

func findSiswa(db *sql.DB, id int) {
	fmt.Println("ID yang dicari :", id)
	var result Siswa
	db.QueryRow("SELECT siswaId,namaSiswa FROM Siswa WHERE siswaId = ?", id).Scan(&result.id, &result.nama)

	fmt.Println(result.id, " ", result.nama)
}

func findSiswaBanyak(db *sql.DB, id int) {
	fmt.Println("ID yang dicari :", id)

	query := "SELECT siswaId,namaSiswa FROM Siswa WHERE siswaId = ?"
	rows, error := db.Query(query, id)

	if error != nil {
		fmt.Println(error.Error())
	}

	var resultAll []Siswa
	for rows.Next() {
		siswa := Siswa{}
		rows.Scan(&siswa.id, &siswa.nama)
		resultAll = append(resultAll, siswa)
	}

	for _, value := range resultAll {
		fmt.Println(value.id, " ", value.nama)
	}
}

func updateSiswa(db *sql.DB, id int, siswa Siswa) {

	var query string

	if siswa.kelas != "" && siswa.nama == "" {
		query = "UPDATE Siswa SET kelas = ? WHERE siswaId = ?"
		row, error := db.Exec(query, siswa.kelas, id)

		if error != nil {
			fmt.Println(error.Error())
		}

		fmt.Print(row.RowsAffected())
		fmt.Println(" affected")
	} else {
		query = "UPDATE Siswa SET namaSiswa = ?,kelas = ? WHERE siswaId = ?"
		row, error := db.Exec(query, siswa.nama, siswa.kelas, id)

		if error != nil {
			fmt.Println(error.Error())
		}

		fmt.Print(row.RowsAffected())
		fmt.Println(" affected")
	}

}

func findSiswaReturn(db *sql.DB, id int) Siswa {
	fmt.Println("ID yang dicari :", id)
	var result Siswa
	db.QueryRow("SELECT siswaId,namaSiswa,kelas FROM Siswa WHERE siswaId = ?", id).Scan(&result.id, &result.nama, &result.kelas)

	return result
}

func updateSiswaCase(db *sql.DB, id int, siswa Siswa) {
	dataSiswa := findSiswaReturn(db, id) // original -> ambil dari db

	if siswa.kelas != "" {
		dataSiswa.kelas = siswa.kelas
	}
	if siswa.nama != "" {
		dataSiswa.nama = siswa.nama
	}

	query := "UPDATE Siswa SET namaSiswa =?,kelas=? WHERE siswaId =?"
	db.Exec(query, dataSiswa.nama, dataSiswa.kelas, id)

	fmt.Println(dataSiswa)
}

func deleteSiswa(db *sql.DB, id int) {
	query := "DELETE FROM Siswa WHERE siswaId = ?"
	row, error := db.Exec(query, id)
	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Print(row.RowsAffected())
	fmt.Println(" affected")
}
