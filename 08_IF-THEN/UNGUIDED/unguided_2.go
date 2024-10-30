package main

import (
	"fmt"
)

func main() {

	// Membuat variabel
	var angka int

	// Inputan untuk memasukan angka
	fmt.Print("Masukan angka: ")
	fmt.Scan(&angka)

	// Percabangan apakah angka genap negatif atau bukan
	if angka < 0 && angka%2 == 0 {
		// Menampilkan output
		fmt.Println("Genap Negatif")
		return
	}

	//Output jika percabangan tidak terpenuhi
	fmt.Println("Bukan")

}
