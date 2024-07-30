package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	// fmt.Print("Hello Dunia \n")
	// fmt.Printf("Selamat World")

	// fmt.Println()

	// var nilai int = 100
	// fmt.Printf("Nilai adalah %d unit \n", nilai)

	// nilai2 := 200
	// fmt.Println("Nilai 2 adalah ", nilai2, " unit")

	// var nilai4, nilai5 uint8 = 100, 200
	// fmt.Println("nilai 4 dan nilai 5 adalah ", nilai4, " ", nilai5)

	// nilai6, nilai7 := 200, 300
	// fmt.Println("nilai 6 dan nilai 7 adalah ", nilai6, " ", nilai7)

	// var data uint16
	// fmt.Print("Masukan nilai anda?")
	// fmt.Scanln(&data)

	// fmt.Println("Data adalah ", data)

	// const NAME string = "CourseNet"
	// fmt.Println("Nilai Nama adalah ", NAME)

	alas := 10
	tinggi := 10
	var luas int
	luas = (alas + tinggi) / 2
	fmt.Println("Luas Segitiga adalah ", luas)

	var alas1 float32 = 10
	var tinggi1 float32 = 5
	var luas1 float32
	luas1 = (alas1 + tinggi1) / 2
	fmt.Printf("Luas Segitiga1 adalah %.2f \n", luas1)

	varA := 132
	varB := varA % 5
	fmt.Println("Nilai VarB adalah ", varB)

}
