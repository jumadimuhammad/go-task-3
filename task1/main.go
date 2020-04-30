package main

import "fmt"

func main() {
	var nama = map[int]string{}
	var harga = map[int]string{}
	var stok = map[int]int{}

	nama[0] = "Baju panjang"
	harga[0] = "Rp. 200.000"
	stok[0] = 8

	nama[1] = "Celana panjang"
	harga[1] = "Rp. 300.000"
	stok[1] = 9

	nama[2] = "Rok mini"
	harga[2] = "Rp. 400.000"
	stok[2] = 19

	nama[3] = "Hoodie"
	harga[3] = "Rp. 500.000"
	stok[3] = 2

	fmt.Println("List Stok Barang")
	fmt.Println("-----------------------------------------")
	for index, value := range stok {
		if value < 10 {
			fmt.Println("Nama barang  : ", nama[index])
			fmt.Println("Harga barang : ", harga[index])
			fmt.Println("Stok barang  : ", stok[index])
			fmt.Println("-----------------------------------------")
		}
	}

}
