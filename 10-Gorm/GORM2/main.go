package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Structs untuk tabel
type Employee struct {
	gorm.Model
	Name      string
	Address   Address
	AddressID uint
}

type Address struct {
	gorm.Model
	Street     string
	City       string
	State      string
	EmployeeID uint
}

type Customer struct {
	gorm.Model
	Name   string
	Orders []Order
}

type Order struct {
	gorm.Model
	OrderNumber string
	CustomerID  uint
}

type Author struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"many2many:author_books"`
}

type Book struct {
	gorm.Model
	Title   string
	Authors []Author `gorm:"many2many:author_books"`
}

func main() {
	// Koneksi ke database MySQL
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrasi tabel
	db.AutoMigrate(&Employee{}, &Address{}, &Customer{}, &Order{}, &Author{}, &Book{})

	// Menambahkan data contoh
	employee := Employee{
		Name: "Alice",
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
			State:  "CA",
		},
	}
	db.Create(&employee)

	customer := Customer{
		Name: "Bob",
		Orders: []Order{
			{OrderNumber: "12345"},
			{OrderNumber: "67890"},
		},
	}
	db.Create(&customer)

	author := Author{
		Name: "Charles",
		Books: []Book{
			{Title: "Golang 101"},
			{Title: "Advanced Golang"},
		},
	}
	db.Create(&author)

	// Membaca data Employee beserta Address
	var readEmployee Employee
	db.Preload("Address").First(&readEmployee, employee.ID)
	fmt.Printf("Employee: %v\n", readEmployee)
	fmt.Printf("Address: %v\n", readEmployee.Address)

	// Membaca data Customer beserta Orders
	var readCustomer Customer
	db.Preload("Orders").First(&readCustomer, customer.ID)
	fmt.Printf("Customer: %v\n", readCustomer)
	fmt.Printf("Orders: %v\n", readCustomer.Orders)

	// Membaca data Author beserta Books
	var readAuthor Author
	db.Preload("Books").First(&readAuthor, author.ID)
	fmt.Printf("Author: %v\n", readAuthor)
	fmt.Printf("Books: %v\n", readAuthor.Books)
}
