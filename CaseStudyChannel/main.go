package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Send Channel
	chSend := make(chan int)
	go sendData(chSend, 42) // Mengirim data ke channel
	received := <-chSend    // Menerima data dari channel
	fmt.Println("Received:", received)

	// Receive Channel
	chReceive := make(chan int)
	go receiveData(chReceive) // Menerima data dari channel
	chReceive <- 42           // Mengirim data ke channel

	// Full Case
	rand.Seed(time.Now().UnixNano())
	// Membuat channel untuk mengirim data pemesanan tiket
	ch := make(chan TicketOrder)
	// Data pemesanan yang diterima oleh setiap kasir
	cashierOrders := [][]TicketOrder{
		{
			{OrderID: 1, Customer: "Alice", Movie: "Movie A"},
			{OrderID: 2, Customer: "Bob", Movie: "Movie B"},
		},
		{
			{OrderID: 3, Customer: "Charlie", Movie: "Movie C"},
			{OrderID: 4, Customer: "Dave", Movie: "Movie D"},
		},
		{
			{OrderID: 5, Customer: "Eve", Movie: "Movie E"},
			{OrderID: 6, Customer: "Frank", Movie: "Movie F"},
		},
	}

	// Menjalankan goroutine untuk setiap kasir
	for i, orders := range cashierOrders {
		go cashier(i+1, orders, ch)
	}

	// Menerima dan mencetak data pemesanan tiket
	for i := 0; i < len(cashierOrders)*2; i++ {
		order := <-ch
		fmt.Printf("Ticket printed for OrderID %d: Customer %s, Movie %s\n", order.OrderID, order.Customer, order.Movie)
	}
}
