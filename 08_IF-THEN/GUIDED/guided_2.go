package main

import (
	"fmt"
)

func main() {

	var angka int

	fmt.Print("Masukan angka: ")
	fmt.Scan(&angka)

	if angka >= 1 {
		fmt.Println("Positif")
	} else {
		fmt.Println("Bukan Positif")
	}
}
