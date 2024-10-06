package main

import (
	"fmt"
)

func main() {

	//Variabel r yang digunakan untuk memuat sebuah jari-jari lingkaran dan phi untuk menyimpan nilai pasti dari operasi lingkaran.
	var r float64
	phi := 3.14

	// Menginputkan sebuah jari-jari
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	//Operasi yang digunakan untuk mencari sebuah luas dan keliling dari lingkaran
	luas := phi * r * r
	keliling := 2 * phi * r

	//Menampilkan hasil dari operasi luas dan keliling lingkaran.
	fmt.Println("Luas lingkaran tersebut yaitu ", luas, " dan keliling lingkaran adalah ", keliling)

}
