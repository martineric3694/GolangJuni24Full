package main

func sendData(ch chan<- int, data int) {
	ch <- data // Mengirim data ke channel
}
