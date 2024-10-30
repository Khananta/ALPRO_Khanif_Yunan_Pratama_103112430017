package main

import (
	"fmt"
)

func main() {

	var nilaiMutlak int

	fmt.Print("Masukan angka: ")
	fmt.Scan(&nilaiMutlak)

	if nilaiMutlak < 0 {
		nilaiMutlak = -1 * nilaiMutlak
	}
	fmt.Println("Nilai mutlaknya adalah ", nilaiMutlak)

}
