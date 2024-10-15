package main

import (
	"fmt"
)

func main() {

	// Membuah sebuah variabel bernama jumlahTitik untuk menyimpan inputan
	var jumlahTitik int

	// Inputan untuk memasukan jumlah baris yang diinginkan
	fmt.Print("Masukkan jumlah baris segitiga: ")
	fmt.Scan(&jumlahTitik)

	// Perulangan untuk membuat segitiga sesuai dengan inputan baris yang sudah dimasukan.
	for i := 1; i <= jumlahTitik; i++ {
		for bintang := 1; bintang <= i; bintang++ {
			// Output berupa perulangan yaitu disimbolkan dengan simbol bintang (*)
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
