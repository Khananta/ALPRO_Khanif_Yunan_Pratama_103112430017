package main

import "fmt"

func main() {

	var (
		masukan, diskon, hasil int
	)

	// Inputan untuk memasukan harga asli dan diskon.
	fmt.Print("Masukan harga: ")
	fmt.Scan(&masukan)

	fmt.Print("Masukan diskon: ")
	fmt.Scan(&diskon)

	hasil = masukan - (masukan * diskon / 100) // Rumus untuk menemukan harga setelah diskon.

	fmt.Print(hasil) // Menampilkan harga setelah diskon.
}
