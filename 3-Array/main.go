package main

import (
	"fmt"
)

func main() {
	name := [3]string{"Course", "Net", "Indonesia"}

	fmt.Println(name)

	// // Cara pertama ambil nilai dari array
	for i := 0; i < len(name); i++ {
		fmt.Println(i, ".", name[i])
	}

	// Cara kedua ambil nilai dari array
	for i, value := range name {
		if i == 1 {
			fmt.Println(i, "", value)
		}
	}

	angka := [3]int{1, 2, 3}

	fmt.Println(angka)

	// data := [2][2]int{}

	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Scanln(&data[i][j])
	// 	}
	// }

	// fmt.Println(data)

	array2 := [...]string{"Hello", "World", "Course", "Net", "Indonesia"}

	fmt.Println("Before : ", len(array2))

	arrayRandom := [3]string{
		2: "Course",
		0: "Net",
		1: "Indonesia",
	}

	fmt.Println(arrayRandom)

}
