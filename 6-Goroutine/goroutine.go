package main

import "fmt"

func print(n int, message string) {
	for i := 0; i < n; i++ {
		fmt.Println(i, ".", message)
	}
}
