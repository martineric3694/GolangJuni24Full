package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	p := Person{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}
	PrintFields(p)
}

func PrintFields(i interface{}) {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		if value == "" {
			value = "kosong"
		}
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}
