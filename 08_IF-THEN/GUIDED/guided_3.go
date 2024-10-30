package main

import (
	"fmt"
)

func main() {
	var angka int
	var hasil bool

	fmt.Print("Masukan angka: ")
	fmt.Scan(&angka)

	if angka < 0 && angka%2 == 0 {
		hasil = true
	}

	fmt.Print(hasil)

}
