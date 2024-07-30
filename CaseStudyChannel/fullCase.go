package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struktur data untuk pemesanan tiket
type TicketOrder struct {
	OrderID  int
	Customer string
	Movie    string
}

// Fungsi untuk mengirim data pemesanan dari kasir ke pusat sistem
func cashier(cashierID int, orders []TicketOrder, ch chan<- TicketOrder) {
	for _, order := range orders {
		// Simulasi waktu pemrosesan pesanan
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Cashier %d processed order: %d\n", cashierID, order.OrderID)
		ch <- order
	}
}
