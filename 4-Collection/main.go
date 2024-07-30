package main

import "fmt"

func main() {
	// name := []string{
	// 	"Course",
	// 	"Net",
	// 	"Indonesia",
	// 	"Raya",
	// }

	// //Cetak Nilai
	// fmt.Println(name)

	// //Cetak Nilai dalam loop
	// for _, value := range name {
	// 	fmt.Println(value)
	// }

	// //Replace Nilai
	// cari := "Net"
	// replace := "Not"
	// index := -1
	// for i, value := range name {
	// 	if value == cari {
	// 		index = i
	// 		break
	// 	}
	// }

	// if index == -1 {
	// 	fmt.Println("Kata ", cari, " tidak ditemukan")
	// } else {
	// 	fmt.Println("Kata ", cari, " ditemukan, diganti dengan kata ", replace)
	// 	name[index] = replace
	// 	fmt.Println("Hasilnya menjadi :")
	// 	for _, value := range name {
	// 		fmt.Println(value)
	// 	}
	// }

	// //Menghapus Nilai
	// cari_hapus := "Not"
	// index_hapus := -1
	// for i, value := range name {
	// 	if value == cari_hapus {
	// 		index_hapus = i
	// 		break
	// 	}
	// }

	// if index_hapus == -1 {
	// 	fmt.Println("Kata ", cari_hapus, " tidak ditemukan")
	// } else {
	// 	fmt.Println("Kata ", cari_hapus, " dihapus")
	// 	fmt.Println(name[:index_hapus], " ", name[index_hapus+1:])
	// 	// name = append(name[:index_hapus], name[index_hapus+1:]...)
	// 	fmt.Println("Hasilnya menjadi :", name)

	// }

	// testMap()
	// sliceMapTest()
	// mapSliceTest()
	var data Student = Student{
		nama:   "Budi",
		umur:   30,
		alamat: "Jakarta",
	}

	var sliceStruct []Student
	data1 := Student{
		nama:   "ABC",
		umur:   100,
		alamat: "Jakata",
	}
	data2 := Student{
		nama:   "NCD",
		umur:   50,
		alamat: "Bandung",
	}

	sliceStruct = append(sliceStruct, data1, data2)

	fmt.Println("Struct Parameter In")
	strucTest(data, sliceStruct)

	kembalian := structReturn(data1.nama, data2.alamat)
	fmt.Println("Struct Return")
	fmt.Println(kembalian)

}
