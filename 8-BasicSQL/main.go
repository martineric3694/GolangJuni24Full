package main

func main() {
	db := connectDB()

	// getAllSiswa(db)

	// siswa := Siswa{
	// 	id:    3,
	// 	nama:  "Susi",
	// 	kelas: "XII",
	// }

	// insertSiswa(db, siswa)

	// getAllSiswa(db)

	// findSiswa(db, 1)

	// findSiswaBanyak(db, 1)
	// siswa := Siswa{
	// 	kelas: "VIII",
	// }
	// updateSiswa(db, 1, siswa)

	// siswa := Siswa{
	// 	nama: "Sandi",
	// 	// kelas: "XII",
	// }

	// updateSiswaCase(db, 1, siswa)

	deleteSiswa(db, 1)
}
