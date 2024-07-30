package main

import (
	"fmt"
)

// Definisi struct Product
type ProductPointer struct {
	name  string
	stock int
}

// Fungsi untuk menambahkan produk baru ke dalam daftar produk
func addProductPointer(products *[]ProductPointer, name string, stock int) {
	*products = append(*products, ProductPointer{name: name, stock: stock})
}

// Fungsi untuk memperbarui stok produk yang ada
func updateStockPointer(products *[]ProductPointer, name string, stock int) {
	for i, product := range *products {
		if product.name == name {
			(*products)[i].stock = stock
			return
		}
	}
	fmt.Println("Produk tidak ditemukan:", name)
}

// Fungsi untuk menampilkan daftar produk beserta jumlah stoknya
func displayProductsPointer(products []ProductPointer) {
	fmt.Println("Daftar Produk dan Stok:")
	for _, product := range products {
		fmt.Printf("Nama: %s, Stok: %d\n", product.name, product.stock)
	}
}
