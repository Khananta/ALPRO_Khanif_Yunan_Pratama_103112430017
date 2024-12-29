package main

import "fmt"

func main() {

	var angka, jumlah, digit int

	fmt.Scan(&angka)

	for angka > 0 {
		digit = angka % 10

		fmt.Print(digit, " ")

		jumlah += digit
		angka /= 10
	}

	fmt.Println("")
	fmt.Println(jumlah)

}
