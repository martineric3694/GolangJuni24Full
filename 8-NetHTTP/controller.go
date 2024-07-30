package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan IP dari Requestor
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Jika remote Address tidak didapat karena menggunakan Load Balancer
		fmt.Println("FORWARED")
		fmt.Printf("IP Address: %s \n", forwarded)
	} else {
		// Default menggunakan Remote Address untuk mendapatkan IP
		fmt.Println("ELSE")
		fmt.Println(r.RemoteAddr)
		fmt.Println(r.RequestURI)
		ipHost := strings.Split(r.RemoteAddr, ":")
		fmt.Printf("IP Address: %s \n", ipHost[0])
	}
	fmt.Fprintln(w, cetakHello())
}

func home(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello ", i)
		fmt.Fprintln(w, "Hello ", i)
	}
}

func maps(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"data 1": "Hello",
		"data 2": 100,
	}

	fmt.Fprintln(w, data)
}

type Data struct {
	Data1 int    `json:"umur"` // tag json untuk mengatur result / output saat mengembalikan nilai, contoh Data1 diubah jadi umur
	Data2 string `json:"nama"`
}

func structs(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ada yang hit ")
	data := Data{
		Data1: 200,
		Data2: "Course Net",
	}
	// cara 1 mengembalikan json dengan cara marshal
	// w.Header().Set("Content-Type", "application/json")
	// result, _ := json.Marshal(data)
	// w.Write(result)

	// cara 2 mengembalikan json dengan cara encode
	json.NewEncoder(w).Encode(&data)
}

func queryString(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.Query()
		data := url.Get("id")
		data1 := url.Get("nama")

		result := map[string]interface{}{
			"statusCode": http.StatusOK,
			"id":         data,
			"name":       data1,
		}
		fmt.Fprintln(w, result)
	} else {
		fmt.Fprintln(w, http.StatusInternalServerError)
	}
}

func postParam(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data Data
		json.NewDecoder(r.Body).Decode(&data)

		fmt.Fprintln(w, "Received data : ", data)
		fmt.Println(data)
	} else {
		fmt.Fprintln(w, http.StatusNotFound)
	}
}

func getAPIData(w http.ResponseWriter, r *http.Request) {
	url := "https://dummyjson.com/recipes"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		fmt.Fprintln(w, http.StatusBadRequest)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error Read Body :", err.Error())
		fmt.Fprintln(w, http.StatusBadRequest)
	}

	// fmt.Println(string(body))

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error Unmarshal : ", err.Error())
		fmt.Fprintln(w, http.StatusBadRequest)
	}

	recipes := data["recipes"].([]interface{})

	for _, v := range recipes {
		fmt.Println(v)
	}

}

func cetakHello() string {
	return "Course Net Gading Serpong Indonesia"
}
