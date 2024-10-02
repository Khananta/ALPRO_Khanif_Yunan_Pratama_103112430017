package main

import (
	"fmt"
	"math"
)

func main() {

	//Membuat wadah untuk menyimpan data sebuah volume, luas dan jari-jari
	var (
		volumeBola float64
		luasBola   float64
		jari       float64
	)

	// Inputan untuk memasukan jari-jari dari sebuah bola
	fmt.Print("Masukan jari-jari: ")
	fmt.Scan(&jari)

	//Operasi untuk menghitung volume dan luas sekaligus menampilkan hasil dari operasi volume dan luas suatu bola.
	volumeBola = (4.0 / 3.0) * math.Pi * math.Pow(jari, 3)
	luasBola = 4 * math.Pi * math.Pow(jari, 2)
	fmt.Printf("Volume bola adalah: %.4f\n", volumeBola)
	fmt.Printf("dan Luas kulit bola adalah: %.4f\n", luasBola)

}
