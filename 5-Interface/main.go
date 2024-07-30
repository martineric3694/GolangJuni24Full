package main

import (
	"fmt"
	"reflect"
)

func main() {

	// fmt.Println(luas(10, 20))
	// fmt.Println(keliling(10))

	// dataSegitiga := Segitiga{
	// 	alas:   20,
	// 	tinggi: 15,
	// }

	// fmt.Println("Segitiga")
	// fmt.Println(dataSegitiga.luas())
	// fmt.Println(dataSegitiga.keliling())

	// dataPersegi := Persegi{
	// 	alas: 10,
	// }

	// fmt.Println("Persegi")
	// fmt.Println(dataPersegi.luas())
	// fmt.Println(dataPersegi.keliling())

	// dataDog := Dog{
	// 	nama: "Anjing Pudel",
	// }

	// dataDog.legs()
	// dataDog.sound()
	// dataDog.breath()

	data := make(map[string]interface{})
	data["nama"] = "Budi"
	data["umur"] = 30
	data["tinggi"] = 170.55

	fmt.Println(reflect.TypeOf(data))

	for _, v := range data {
		fmt.Println(v, " tipe data ", reflect.TypeOf(v))
	}

	data2 := map[string]interface{}{
		"nama":   "Adi",
		"alamat": "Jakarta",
		"bb":     100,
	}

	for _, v := range data2 {
		fmt.Println(v, " tipe data ", reflect.TypeOf(v))
	}
}
