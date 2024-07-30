package main

import "fmt"

type Student struct {
	nama   string
	umur   int
	alamat string
}
type Kelas struct {
	kelas     int
	deskripsi string
	siswa     []Student
}

func strucTest(data Student, sliceStruct []Student) {

	fmt.Println(data)

	fmt.Println(sliceStruct)

}

func structReturn(nama, alamat string) Student {
	data := Student{
		nama:   nama,
		alamat: alamat,
	}

	return data
}
