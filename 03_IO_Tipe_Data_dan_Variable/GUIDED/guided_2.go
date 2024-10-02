package main

import (
	"fmt"
)

func main() {

	var (
		alas   float32
		tinggi float32
		hasil  float32
	)

	fmt.Print("Masukan panjang alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukan tinggi alas segitiga: ")
	fmt.Scan(&tinggi)

	hasil = (alas * tinggi) * 1 / 2 //Rumus untuk menentukan luas segitiga
	fmt.Println("Luas segitiga adalah: ", hasil)

}
