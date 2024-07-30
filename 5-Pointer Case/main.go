package main

func main() {
	var productsPointer []ProductPointer

	// Menambahkan produk baru
	addProductPointer(&productsPointer, "Beras", 50)
	addProductPointer(&productsPointer, "Gula", 30)
	addProductPointer(&productsPointer, "Minyak Goreng", 20)

	// Menampilkan daftar produk awal
	displayProductsPointer(productsPointer)

	// Memperbarui stok produk
	updateStockPointer(&productsPointer, "Gula", 40)
	updateStockPointer(&productsPointer, "Minyak Goreng", 25)

	// Menampilkan daftar produk setelah update stok
	displayProductsPointer(productsPointer)

	var products []Product

	// Menambahkan produk baru
	products = addProduct(products, "Beras", 50)
	products = addProduct(products, "Gula", 30)
	products = addProduct(products, "Minyak Goreng", 20)

	// Menampilkan daftar produk awal
	displayProducts(products)

	// Memperbarui stok produk
	products = updateStock(products, "Gula", 40)
	products = updateStock(products, "Minyak Goreng", 25)

	// Menampilkan daftar produk setelah update stok
	displayProducts(products)
}
