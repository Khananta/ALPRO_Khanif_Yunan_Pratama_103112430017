package main

import "fmt"

func main() {

	// Variabel yang digunakan untuk menyimpan data
	var (
		masukan, dinar, dirham, fals, qirat int
	)

	// Inputan berupa mata uang qirat
	fmt.Print("Masukan jumlah qirat: ")
	fmt.Scan(&masukan)

	// Operasi yang digunakan setiap mata uang
	dinar = masukan / 600
	dirham = (masukan % 600) / 60
	fals = (masukan / 6) % 10
	qirat = masukan % 6

	// Menampilkan hasil dari konversi nilai mata uang
	fmt.Println(dinar, dirham, fals, qirat)

}
