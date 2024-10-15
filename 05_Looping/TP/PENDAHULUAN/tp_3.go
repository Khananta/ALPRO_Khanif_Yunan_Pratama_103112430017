package main

import "fmt"

func main() {

	// Variabel untuk menyimpan range angka
	var (
		rangeAngka int
	)

	// Inputan untuk memasukan range angka
	fmt.Print("Masukkan range angka: ")
	fmt.Scan(&rangeAngka)

	// Menampilkan bilangan genap dari angka yang dimasukan
	for i := 1; i <= rangeAngka; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
