package main

import "fmt"

func main() {

	// Membuat sebuah variabel x dan y
	var (
		nilaiN int
	)

	// Menginputkan nilai dari N
	fmt.Print("Masukan nilai N: ")
	fmt.Scan(&nilaiN)

	// Perulangan untuk menampilkan hasil dari kuadrat (1^2, 2^2, sampai dengan N^2)
	for i := 1; i <= nilaiN; i++ {
		fmt.Print(i*i, " ")
	}
}
