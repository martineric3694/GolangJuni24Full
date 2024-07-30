package main

import (
	"fmt"
	"regexp"
)

// Fungsi untuk memvalidasi email menggunakan regex
func isValidEmail(email string) bool {
	// Pola regex untuk email
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Kompilasi pola regex
	re := regexp.MustCompile(emailPattern)

	// Cocokkan string email dengan pola regex
	return re.MatchString(email)
}

// Validasi Angka
func isNumber(data string) bool {
	const numberPattern = `^\d+$`

	re := regexp.MustCompile(numberPattern)

	return re.MatchString(data)
}

// Validasi Huruf
func isLetter(data string) bool {
	const letterPattern = `^[a-zA-Z]+$`

	re := regexp.MustCompile(letterPattern)

	return re.MatchString(data)
}

func main() {
	// Contoh email untuk diuji
	emails := []string{
		"test@example.com",
		"invalid-email.com",
		"another.test@domain.co",
		"user@sub.domain.com",
		"wrong@domain",
		"a_@abc.com",
	}

	// Periksa setiap email apakah valid atau tidak
	for _, email := range emails {
		if isValidEmail(email) {
			fmt.Printf("Email '%s' adalah valid.\n", email)
		} else {
			fmt.Printf("Email '%s' adalah tidak valid.\n", email)
		}
	}

	dataNumber := []string{
		"123456",
		"123asd456",
		"q123qw4123213",
		"a",
	}
	for _, v := range dataNumber {
		if isNumber(v) {
			fmt.Println(v, " is valid")
		} else {
			fmt.Println(v, " is not valid")
		}
	}

	dataLetter := []string{
		"ASDASDASD",
		"asdasdasdasd",
		"asd12easdasd",
	}

	for _, v := range dataLetter {
		if isLetter(v) {
			fmt.Println(v, " is valid")
		} else {
			fmt.Println(v, " is not valid")
		}
	}
}
