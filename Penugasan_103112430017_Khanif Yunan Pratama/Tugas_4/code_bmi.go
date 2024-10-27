package main

import "fmt"

func main() {

	var (
		nama                           string
		beratBadan, tinggiBadan, hasil float32
	)

	fmt.Print("Masukan nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukan berat badan: ")
	fmt.Scan(&beratBadan)

	fmt.Print("Masukan tinggi badan: ")
	fmt.Scan(&tinggiBadan)

	hasil = beratBadan / (tinggiBadan * tinggiBadan)

	fmt.Print("Informasi BMI: ")
	fmt.Println("Nama: ", nama)
	fmt.Printf("Berat: %.2f\n", beratBadan)
	fmt.Printf("Tinggi: %.2f\n", tinggiBadan)
	fmt.Printf("BMI: %.2f\n", hasil)
}
