package main

import (
	"fmt"
)

func main() {

	var (
		r float64
		l float64
		k float64
	)

	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	l = 3.14 * r * r
	k = 2 * 3.14 * r

	fmt.Println("Hasil dari luas lingkaran dengan jari-jari:", r, "adalah", l)
	fmt.Println("Hasil dari keliling lingkaran dengan jari-jari:", r, "adalah", k)
}
