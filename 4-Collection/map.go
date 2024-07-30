package main

import (
	"fmt"
)

func testMap() {

	var data map[string]string = map[string]string{
		"var1": "Course",
		"var2": "Net",
		"var3": "Indonesia",
		"var4": "Gading Serpong",
		"var5": "Indonesia",
		"var6": "Net",
	}

	fmt.Println("Data :", data["var2"])

	for key, value := range data {
		fmt.Println("Key : ", key, " valuenya adalah ", value)
	}

	cari := "Indonesia"
	result := []string{}
	for key, value := range data {
		if value == cari {
			result = append(result, key)
		}
	}

	if result == nil {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("Kata ", cari, " ditemukan pada key : ", result)
	}

	for key, value := range data {
		if value == cari {
			delete(data, key)
			fmt.Println("Kata ", cari, " pada key ", key, " sudah terhapus")
		}
	}

	fmt.Println(data)
}

func sliceMapTest() {
	// data := []map[string]string{}
	var data []map[string]string

	detail := map[string]string{
		"nama":   "Budi",
		"umur":   "30",
		"alamat": "Jakarta",
	}
	detail1 := map[string]string{
		"nama": "Adi",
		"umur": "25",
	}
	detail2 := map[string]string{
		"nama": "Adi1",
		"umur": "251",
	}

	data = append(data, detail, detail1, detail2)
	// Cara 1
	// for i, value := range data {
	// 	if i == 1 {
	// 		continue
	// 	}
	// 	fmt.Println("Nama : ", value["nama"])
	// 	fmt.Println("Umur : ", value["umur"])
	// 	fmt.Println("Alamat :", len(value["alamat"]))
	// }

	// Cara 2
	for _, value := range data { //slice
		for key, valueIn := range value { //map
			if key != "nama" {
				fmt.Println(key, " ", valueIn)
			}
		}
	}
}

func mapSliceTest() {
	data := map[string][]string{
		"detail1": {
			"Training",
			"Center",
		},
		"detail2": {
			"Course",
			"Net",
		},
	}

	for key, value := range data { //map
		for i, valueIn := range value { // slice
			fmt.Println("(", key, ")", i, ".", valueIn)
		}
	}

}
