package main

import "fmt"

func main() {

	var curahHujan string

	fmt.Scan(&curahHujan)

	if curahHujan == "sangattinggi" {
		fmt.Print("macet")
	} else if curahHujan == "tinggi" || curahHujan == "rendah" {
		fmt.Print("tidak macet")
	} else {
		fmt.Print("masukan belum tepat")
	}
}
