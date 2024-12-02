package main

import "fmt"

func main() {

	var kata string

	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	// Perulangan yang akan terus berjalan dan akan berhenti ketika memasukkan kata "telkom"
	for kata != "telkom" {
		fmt.Println("Anda mengetik:", kata)

		fmt.Print("Masukkan kata: ")
		fmt.Scan(&kata)
	}

	fmt.Println("Program Selesai")

}
