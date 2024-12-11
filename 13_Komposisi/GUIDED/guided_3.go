package main

import "fmt"

func main() {

	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Perulangan untuk menampilkan angka dari 1 hingga angka yang dimasukan
	for i := 1; i <= angka; i++ {
		// Percabangan untuk menentukan mana saja faktor dari bilangan tersebut
		if angka%i == 0 {
			fmt.Print(i, " ")
		}
	}

}
