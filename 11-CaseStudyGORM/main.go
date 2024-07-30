package main

import "CaseStudyGorm/controller"

func main() {
	controller.Initiate()

	// controller.InsertKelas()
	// controller.InsertPelajaran()
	// controller.InsertSiswa()
	// controller.InsertNilai()

	// controller.ShowSiswa()
	// controller.ShowKelas()
	// controller.ShowNilaiBySiswa()
	controller.ShowAverageNilaiSiswa()
}
