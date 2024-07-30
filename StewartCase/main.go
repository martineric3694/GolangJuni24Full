package main

import (
	"fmt"
	"reflect"
)

type Siswa struct {
	idSiswa   int
	namaSiswa string
}

func main() {
	siswa := Siswa{
		idSiswa:   1,
		namaSiswa: "Budi",
	}

	valSiswa := reflect.ValueOf(siswa)
	typSiswa := reflect.TypeOf(siswa)

	for i := 0; i < valSiswa.NumField(); i++ {
		key := typSiswa.Field(i).Name
		val := valSiswa.Field(i)

		if val.Kind() == reflect.String {
			if val.String() == "" {
				fmt.Println(key, " kosong")
			}
		} else {
			fmt.Println(key, " ", val)
		}

	}
}
