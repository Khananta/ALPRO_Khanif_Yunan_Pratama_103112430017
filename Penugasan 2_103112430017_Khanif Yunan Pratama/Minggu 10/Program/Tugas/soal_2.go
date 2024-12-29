package main

import "fmt"

func main() {

	var angka1, angka2, angka3, angka4, angka5 float32

	fmt.Scan(&angka1, &angka2, &angka3, &angka4, &angka5)

	if angka1 > angka2 && angka2 > angka3 && angka3 > angka4 && angka4 > angka5 || angka1 < angka2 && angka2 < angka3 && angka3 < angka4 && angka4 < angka5 {
		fmt.Println("Stabil naik/turun")
	} else {
		fmt.Println("Tidak stabil")
	}

}
