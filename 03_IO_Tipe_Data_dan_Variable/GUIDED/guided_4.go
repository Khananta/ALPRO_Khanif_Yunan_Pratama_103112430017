package main

import (
	"fmt"
)

func main() {

	var rupiah int
	var dolar int

	fmt.Print("Masukan nilai Rupiah: ")
	fmt.Scan(&rupiah)

	dolar = rupiah / 15000
	fmt.Println("Konfersi Rupiah ke Dolar adalah: ", dolar)

}
