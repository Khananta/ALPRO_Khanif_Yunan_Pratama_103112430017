package main

import (
	"fmt"
)

func main() {

	var sisi int
	var hasil int

	fmt.Print("Masukan panjang sisi: ")
	fmt.Scan(&sisi)

	hasil = sisi * sisi * sisi
	fmt.Println("Volume kubus dari sisi tersebut adalah: ", hasil)

}
