package main

import (
	"fmt"
	"net/http"
)

const PORT string = ":8080"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/map", maps)
	http.HandleFunc("/struct", structs)
	http.HandleFunc("/postParam", postParam)
	http.HandleFunc("/queryString", queryString)
	http.HandleFunc("/getAPI", getAPIData)
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(PORT, nil)
}
