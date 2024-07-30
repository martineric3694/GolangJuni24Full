package main

import "fmt"

var message = make(chan []string)

var works = make(chan string)

func main() {

	//Case 1 Channel
	go printChannel(5, "Course Net")
	go printChannel(3, "Indonesia")
	go printChannel(7, "Gading Serpong")

	var message1 = <-message
	fmt.Println(message1)

	var message2 = <-message
	fmt.Println(message2)

	var message3 = <-message
	fmt.Println(message3)

	//Case 2 Channel
	// n := 10
	// for i := 0; i < n; i++ {
	// 	go worker(i)
	// }

	// for i := 0; i < n; i++ {
	// 	result := <-works
	// 	if result == "2" {
	// 		fmt.Println("tidak usah dibayar")
	// 	} else {
	// 		fmt.Println(result)
	// 	}
	// }
}
