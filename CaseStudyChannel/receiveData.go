package main

import (
	"fmt"
)

func receiveData(ch <-chan int) {
	data := <-ch // Menerima data dari channel
	fmt.Println("Received:", data)
}
