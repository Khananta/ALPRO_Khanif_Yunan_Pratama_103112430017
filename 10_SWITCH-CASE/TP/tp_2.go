package main

import "fmt"

func main() {

	// Membuat variable
	var umur int

	// Inputan untuk memasukan umur
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)

	// Switch case untuk menentukan kategori menurut umur atau usia
	// Kurang lebih sama seperti penggunaan If else
	switch {
	case umur >= 0 && umur <= 12:
		fmt.Println("Kategori: Anak-anak")
	case umur >= 13 && umur <= 17:
		fmt.Println("Kategori: Remaja")
	case umur >= 18 && umur <= 64:
		fmt.Println("Kategori: Dewasa")
	default:
		fmt.Println("Kategori: Lansia")
	}

}
