package main

import "fmt"

func main() {

	// Membuat variable
	var a, b, hasil int

	// Inputan untuk memasukan angka
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&a)

	// Switch case untuk menentukan kategori dari angka yang dimamsukan
	switch {
	// Digunakan untuk menentukan apakah kelipatan 10
	case a%10 == 0:
		hasil = a / 10

		fmt.Println("Kategori: Kelipatan 10")
		fmt.Println("Hasil pembagian dengan bilangan berikutnya ", a, "/", "10 =", hasil)
	// Menentukan apakah kelipatan 5
	case a%5 == 0 && a > 5:
		hasil = a * a

		fmt.Println("Kategori: Kelipatan 5")
		fmt.Println("Hasil perkalian dengan bilangan berikutnya ", a, "^2", " = ", hasil)
	// Menentukan apakah genap
	case a%2 == 0:
		b = a + 1
		hasil = a * b

		fmt.Println("Kategori: Genap")
		fmt.Println("Hasil perkalian dengan bilangan berikutnya ", a, "*", b, " = ", hasil)
	// Menentukan apakah ganjil
	default:
		b = a + 1
		hasil = a + b

		fmt.Println("Kategori: Ganjil")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya ", a, "+", b, " = ", hasil)
	}

}
