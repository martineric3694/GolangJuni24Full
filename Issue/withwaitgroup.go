package main

import (
	"fmt"
	"sync"
)

func withWaitGroup() {

	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik", "bumbu", "buah", "sterofoam", "bawang", "ikan", "ikan", "daging", "plastik", "sayur", "ikan", "daging", "plastik", "sayur", "bawang",
		"buah", "storofoam", "bawang", "daging", "plastik", "bawang", "buah", "sayur", "bawang", "buah", "sterofoam", "ikan",
	}

	var wg sync.WaitGroup

	i := 1

	for _, brg := range belanja {
		fmt.Print(i)
		fmt.Print(" ")
		switch brg {
		case "ikan":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go cuciwg(brg, &wg)
			i++
		case "daging":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go potongwg(brg, &wg)
			i++
		case "sayur":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go kupaswg(brg, &wg)
			i++
		case "buah":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go cuciwg(brg, &wg)
			i++
		case "bawang":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go kupaswg(brg, &wg)
			i++
		default:
			wg.Add(1)
			go buangSampahwg(brg, &wg)
			i++
		}
	}
	wg.Wait()
}

func cuciwg(barang string, wg *sync.WaitGroup) {
	fmt.Println("Cuci", barang)
	if barang == "ikan" {
		simpanFrozenwg(barang, wg)
	} else {
		simpanChillerwg(barang, wg)
	}
}

func potongwg(barang string, wg *sync.WaitGroup) {
	fmt.Println("Potong", barang)
	simpanFrozenwg(barang, wg)
}

func kupaswg(barang string, wg *sync.WaitGroup) {
	fmt.Println("Cuci", barang)
	simpanBoxwg(barang, wg)
	// simpanChillerwg(barang, wg)
}

func simpanBoxwg(barang string, wg *sync.WaitGroup) {
	fmt.Println(barang, "Simpan Box")
	simpanChillerwg(barang, wg)
}

func simpanFrozenwg(barang string, wg *sync.WaitGroup) {
	fmt.Println(barang, "Simpan Frozen")
	defer wg.Done()
}

func simpanChillerwg(barang string, wg *sync.WaitGroup) {
	fmt.Println(barang, "Simpan Chiller")
	defer wg.Done()
}

func buangSampahwg(barang string, wg *sync.WaitGroup) {
	fmt.Println(barang, "Buang Sampah")
	defer wg.Done()
}
