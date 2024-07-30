package main

import (
	"fmt"
	"strconv"
)

func casting() {

	// Casting Int dengan Int lainnya
	var intVal int = 42
	var int8Val int8 = int8(intVal)
	var int16Val int16 = int16(intVal)
	var int32Val int32 = int32(intVal)
	var int64Val int64 = int64(intVal)

	fmt.Println(int8Val, int16Val, int32Val, int64Val)

	// Casting Float dengan Float lainnya
	var float32Val float32 = 42.42
	var float64Val float64 = float64(float32Val)

	fmt.Println(float32Val, float64Val)

	// Casting int ke float dan sebaliknya
	var intVal_1 int = 42
	var float64Val_1 float64 = float64(intVal_1 + intVal_1)

	var floatVal float64 = 42.42
	var intValFromFloat int = int(floatVal)

	fmt.Println(float64Val_1, intValFromFloat)

	// Casting string ke slice byte
	var strVal string = "Hello, World!"
	var byteSlice []byte = []byte(strVal)
	var newStr string = string(byteSlice)

	fmt.Println(byteSlice, newStr)

	// Konversi dari string ke int
	intVal, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intVal)

	// Konversi dari int ke string
	strVal_1 := strconv.Itoa(42)
	fmt.Println(strVal_1)

	// Konversi dari string ke float
	floatVal_1, err := strconv.ParseFloat("42.42", 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(floatVal_1)

	// Konversi dari float ke string
	strFloatVal := strconv.FormatFloat(42.42, 'f', 2, 64)
	fmt.Println(strFloatVal)
}
