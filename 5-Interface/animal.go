package main

import "fmt"

type Animal interface {
	legs()
	sound()
	breath()
}

type Dog struct {
	nama string
}

func (d Dog) legs() {
	fmt.Println(d.nama, " memiliki 4 kaki")
}
func (d Dog) sound() {
	fmt.Println(d.nama, " menggonggong")
}
func (d Dog) breath() {
	fmt.Println(d.nama, " bernafas menggunakan paru-paru")
}
