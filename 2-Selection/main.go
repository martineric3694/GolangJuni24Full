package main

import (
	"fmt"
	"strconv"
)

func main() {
	varA := 100
	varB := 200

	fmt.Println("Apakah nilai varA dan varB sama?", varA >= varB)

	if varA >= varB {
		if varA == 100 {

		}
		varA += 100
	} else {
		varB -= 100
	}
	fmt.Printf("VarA : %d, VarB :%d\n", varA, varB)

	data := fmt.Sprintf("nilai varA adalah %d", varA) //format string
	fmt.Println(data)

	funcGet := getData(100)
	fmt.Println(funcGet)

	angka := "100a"
	varAngka, error := strconv.Atoi(angka)
	if error != nil {
		fmt.Println("ini bukan angka")
	} else {
		fmt.Printf("Angka adalah %d\n", varAngka)
	}

	switch varAngka {
	case 100:
		fmt.Println("Nilai nya adalah ", varAngka)
	case 200:
		fmt.Println("Nilai nya adalah ", varAngka)
	default:
		fmt.Println("Tidak ada nilainya")
	}

	if varAngka == 100 {
		fmt.Println("Nilai nya adalah ", varAngka)
	} else if varAngka == 200 {
		fmt.Println("Nilai nya adalah ", varAngka)
	} else {
		fmt.Println("Tidak ada nilainya")
	}
}

func getData(n int) string {
	data := fmt.Sprintf("Nilai dari N adalah %d", n)
	return data
}
