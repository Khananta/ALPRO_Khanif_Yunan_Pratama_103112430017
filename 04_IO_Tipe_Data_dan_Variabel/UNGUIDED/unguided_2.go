package main

import (
	"fmt"
)

func main() {

	var (
		bmi, tinggiBadan, hasil float32
	)

	// Masukan atau inputan untuk memasukan BMI dan tinggi badan
	fmt.Print("Masukan BMI: ")
	fmt.Scan(&bmi)

	fmt.Print("Masukan tinggi badan: ")
	fmt.Scan(&tinggiBadan)

	hasil = bmi * (tinggiBadan * tinggiBadan) // Rumus untuk mencari sebuah berat badan dari bmi dan tinggi badan.

	fmt.Printf("Berat Badan adalah %.0f\n", hasil) // Menampilkan hasil dari operasi.
}
