package main

import (
	"fmt"
)

func main() {

	// Membuat variable
	var kode int

	// Menampilakn daftar menu yang ada
	fmt.Println("Daftar Menu Restoran:")
	fmt.Println("1. Burger - Rp25.000")
	fmt.Println("2. Fried Chicken - Rp20.000")
	fmt.Println("3. French Fries - Rp15.000")
	fmt.Println("4. Soft Drink - Rp10.000")
	fmt.Println("5. Coffee - Rp15.000")

	// Meminta untuk memasukan kode makanan sesuai dengan menu
	fmt.Print("Masukkan kode item (1-5): ")
	fmt.Scan(&kode)

	// Switch case untuk memilih makanan sesuai dengan kode yang sudah dimasukan.
	switch kode {
	case 1:
		fmt.Println("Anda memilih: Burger - Rp25.000")
	case 2:
		fmt.Println("Anda memilih: Fried Chicken - Rp20.000")
	case 3:
		fmt.Println("Anda memilih: French Fries - Rp15.000")
	case 4:
		fmt.Println("Anda memilih: Soft Drink - Rp10.000")
	case 5:
		fmt.Println("Anda memilih: Coffee - Rp15.000")
	// Default digunakan untuk menampilkan teks jika kode yang dimasukan tidak sesuai
	default:
		fmt.Println("Kode menu tidak valid")
	}
}
