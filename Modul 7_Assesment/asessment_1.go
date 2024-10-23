package main

import (
	"fmt"
)

func main() {
	// Membuat sebuah variabel x dan y
	var (
		x, y, hasil int
	)

	// Menginputkan nilai x dan y
	fmt.Print("Masukan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukan nilai y: ")
	fmt.Scan(&y)

	hasil = 0
	// Menjumlahkan perulangan dari x sampai y
	for i := x; i <= y; i++ {
		hasil = hasil + i
	}

	// Menampilkan hasil dari x sampai y
	fmt.Println("Hasil dari", x, "sampai", y, "adalah", hasil)
}
