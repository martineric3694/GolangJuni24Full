package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	time.Sleep(time.Duration(id) * time.Second)
	works <- fmt.Sprint("Workers ", id, " menyelesaikan pekerjaan ", time.Duration(id))
}
