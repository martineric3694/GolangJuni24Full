package main

import (
	"time"
)

func main() {
	// Goroutine

	go print(100, "Course Net")
	go print(200, "Indonesia")
	go print(150, "Gading Serpong")

	// fmt.Scanln()
	time.Sleep(5 * time.Second)

	// Non Goroutine
	// print(100, "Course Net")
	// print(200, "Indonesia")
	// print(150, "Gading Serpong")

}
