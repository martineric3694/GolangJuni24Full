package main

import "net/http"

func main() {
	http.HandleFunc("/", getAll)
	http.HandleFunc("/getOne", getOne)
	http.ListenAndServe(":8080", nil)
}
