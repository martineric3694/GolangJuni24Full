package main

import (
	"fmt"
	"sync"
)

func printWG(n int, message string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		fmt.Println(i, ".", message)
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mengurangi counter WaitGroup saat fungsi selesai
	fmt.Printf("Worker %d starting\n", id)
	fmt.Printf("Worker %d done\n", id)
}
