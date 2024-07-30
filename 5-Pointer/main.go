package main

import (
	"fmt"
	"reflect"
)

type Siswa struct {
	nama string
	umur int
}

func main() {
	a := 10
	ptr := &a
	var b int
	b = a

	// Dengan pointer
	fmt.Println("Dengan Pointer")
	fmt.Println(a)
	*ptr = 20
	fmt.Println(a)

	// Tanpa pointer
	fmt.Println("Tanpa Pointer")
	fmt.Println(b)
	a = 50
	fmt.Println(b)
	fmt.Println(ptr)

	data := Siswa{
		nama: "Budi",
		umur: 30,
	}

	fmt.Println(reflect.TypeOf(data))
}
