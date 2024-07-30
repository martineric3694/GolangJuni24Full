package main

import (
	"fmt"
)

// Definisi struct Product
type Product struct {
	name  string
	stock int
}

// Fungsi untuk menambahkan produk baru ke dalam daftar produk
func addProduct(products []Product, name string, stock int) []Product {
	products = append(products, Product{name: name, stock: stock})
	return products
}

// Fungsi untuk memperbarui stok produk yang ada
func updateStock(products []Product, name string, stock int) []Product {
	for i, product := range products {
		if product.name == name {
			products[i].stock = stock
			return products
		}
	}
	fmt.Println("Produk tidak ditemukan:", name)
	return products
}

// Fungsi untuk menampilkan daftar produk beserta jumlah stoknya
func displayProducts(products []Product) {
	fmt.Println("Daftar Produk dan Stok:")
	for _, product := range products {
		fmt.Printf("Nama: %s, Stok: %d\n", product.name, product.stock)
	}
}
