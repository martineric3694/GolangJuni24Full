package main

type Siswa struct {
	idSiswa   int
	namaSiswa string
	kelas     string
}

func getAllDB() []Siswa {
	db := connectDB()
	var result []Siswa
	// call query
	defer db.Close()
	return result
}
