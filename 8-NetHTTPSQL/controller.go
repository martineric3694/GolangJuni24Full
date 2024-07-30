package main

import (
	"fmt"
	"net/http"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getAllDB())
}

func getOne(w http.ResponseWriter, r *http.Request) {

}
