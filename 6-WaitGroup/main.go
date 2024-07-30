package main

import (
	"fmt"
	"sync"
)

func main() {

	// var wg sync.WaitGroup

	// wg.Add(1)
	// go printWG(100, "Course Net", &wg)

	// wg.Add(1)
	// go printWG(200, "Indonesia", &wg)

	// wg.Add(1)
	// go printWG(150, "Gading Serpong", &wg)

	// wg.Wait()
	// fmt.Println("All task done")

	var wg sync.WaitGroup

	for i := 1; i <= 25; i++ {
		wg.Add(1) // menambah counter WaitGroup
		go worker(i, &wg)
	}

	wg.Wait() // menunggu semua goroutine selesai
	fmt.Println("All workers done")
}
