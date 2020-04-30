package main

import "fmt"

func main() {
	var produk1 = map[string]interface{}{
		"namaBarang":  "Baju panjang",
		"hargaBarang": "Rp 200.000",
		"stokBarang":  9,
	}

	var produk2 = map[string]interface{}{
		"namaBarang":  "Celana jeans",
		"hargaBarang": "Rp 400.000",
		"stokBarang":  10,
	}

	var produk3 = map[string]interface{}{
		"namaBarang":  "Rok mini",
		"hargaBarang": "Rp 400.000",
		"stokBarang":  4,
	}

	var produk4 = map[string]interface{}{
		"namaBarang":  "Hoddie",
		"hargaBarang": "Rp 500.000",
		"stokBarang":  14,
	}

	if produk1["stokBarang"].(int) < 10 {
		fmt.Println("Nama barang : ", produk1["namaBarang"].(string))
		fmt.Println("Harga barang : ", produk1["hargaBarang"].(string))
		fmt.Println("Stok barang : ", produk1["stokBarang"].(int))
	}

	if produk2["stokBarang"].(int) < 10 {
		fmt.Println("Nama barang : ", produk2["namaBarang"].(string))
		fmt.Println("Harga barang : ", produk2["hargaBarang"].(string))
		fmt.Println("Stok barang : ", produk2["stokBarang"].(int))
	}

	if produk3["stokBarang"].(int) < 10 {
		fmt.Println("Nama barang : ", produk3["namaBarang"].(string))
		fmt.Println("Harga barang : ", produk3["hargaBarang"].(string))
		fmt.Println("Stok barang : ", produk3["stokBarang"].(int))
	}

	if produk4["stokBarang"].(int) < 10 {
		fmt.Println("Nama barang : ", produk4["namaBarang"].(string))
		fmt.Println("Harga barang : ", produk4["hargaBarang"].(string))
		fmt.Println("Stok barang : ", produk4["stokBarang"].(int))
	}

	var product = []map[string]interface{}{
		{"namaProduct": "Baju panjang", "hargaProduct": "Rp. 300.000", "stockProduct": 8},
		{"namaProduct": "Celana Jeans", "hargaProduct": "Rp. 400.000", "stockProduct": 9},
		{"namaProduct": "Hoodie", "hargaProduct": "Rp. 300.000", "stockProduct": 12},
		{"namaProduct": "Rok mini", "hargaProduct": "Rp. 100.000", "stockProduct": 2},
	}

	for _, item := range product {
		if item["stockProduct"].(int) < 10 {
			fmt.Println("Nama produk :", item["namaProduct"])
			fmt.Println("Harga produk :", item["hargaProduct"])
			fmt.Println("Stok produk :", item["stockProduct"])
			fmt.Println("------------------------------------------")
		}
	}

}