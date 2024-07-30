package main

import "fmt"

func printChannel(n int, data string) {
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, fmt.Sprintln(data, "(", i, ")"))
	}
	message <- result
}
