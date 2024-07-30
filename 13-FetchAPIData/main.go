package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.GET("/data", GetData)
	route.Run()
}

func GetData(c *gin.Context) {

	url := "https://dummy.restapiexample.com/api/v1/employees"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var body_json Body
	json.Unmarshal(body, &body_json)

	for _, v := range body_json.Data {
		fmt.Println(v.Id)
	}

	c.JSON(200, gin.H{
		"body_json": body_json.Data,
	})
}

type Body struct {
	Status  string `json:"status"`
	Data    []Data `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	Id      uint   `json:"id"`
	Name    string `json:"employee_name"`
	Salary  int    `json:"employee_salary"`
	Age     uint   `json:"employee_age"`
	Profile string `json:"profile_image"`
}
