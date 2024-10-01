package main

import (
	"fmt"
)

func main() {

	var (
		nama  string
		nim   int
		kelas string
	)

	fmt.Print("Masukan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan NIM: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukan Kelas: ")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalkan saya adalah", nama, "salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM", nim)
}
