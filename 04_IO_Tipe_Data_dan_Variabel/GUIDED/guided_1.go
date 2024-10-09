package main

import (
	"fmt"
)

func main() {

	var (
		jam, menit, detik, masukan int
	)

	fmt.Print("Masukan detik: ")
	fmt.Scan(&masukan)

	jam = masukan / 3600
	menit = (masukan % 3600) / 60
	detik = (masukan % 3600) % 60

	fmt.Print(jam, " Jam ", menit, " Menit ", detik, " Detik")

}
