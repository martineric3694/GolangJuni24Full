package main

import "fmt"

func main() {
	// data := []string{}

	// fmt.Println("Before : ", len(data), " Data :", data)

	// data = append(data, "Hellow", "World", "Hellow1", "World1", "Hellow2", "World2")

	// fmt.Println("After : ", len(data), " Data :", data)
	// fmt.Println("Data ke 0 adalah :", data[0])

	// var cari string = "Hellow"
	// var ganti string = "Hello"
	// data[0] = "CourseNet"

	// fmt.Println(data[6:])

	// for _, value := range data {
	// 	if value == cari {
	// 		value = ganti
	// 		fmt.Println(value)
	// 	}
	// }
	makeSlice()
}

func makeSlice() {
	// Membuat slice dengan panjang 5 dan kapasitas 10
	s := make([]string, 6, 10)

	// Mengisi nilai slice
	for i := 0; i < 5; i++ {
		s = append(s, "Hello")
	}

	// Menambahkan elemen tambahan
	s = append(s, "World", "Course", "Net")

	fmt.Println("Slice:", s)
	fmt.Println("Panjang:", len(s), "Kapasitas:", cap(s))
}

func defaultSlice() {
	data := []string{}
	for i := 0; i < 5; i++ {
		data = append(data, "Hellow")
	}

	fmt.Println("Slice:", data)
	fmt.Println("Panjang:", len(data), "Kapasitas:", cap(data))
}

// Buatkan search engine
// 1. Menambahkan kosakata / kamus
// 2. Mencari kata di dalam kosakata / kamus
// 3. Mengganti nilai kata yang dicari menjadi kata yang diinginkan
// 4. Menampilkan semua kosakata
