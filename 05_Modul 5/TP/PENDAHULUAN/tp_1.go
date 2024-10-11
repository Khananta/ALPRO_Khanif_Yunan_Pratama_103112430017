package main

import (
	"fmt"
)

func main() {

	// Membuah sebuah variabel bernama nilaiN untuk menyimpan inputan.
	var (
		nilaiN, hitung int
	)

	// Inputan untuk memasukan bilangan bulat.
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&nilaiN)

	// Rumus untuk mencari jumlah deret angka menggunakan perulangan
	hitung = 0
	for i := 1; i <= nilaiN; i++ {
		hitung += i
	}

	// Output atau hasil dari operasi diatas
	fmt.Print("Jumlah dari deret 1 hingga ", nilaiN, " adalah ", hitung)
}
