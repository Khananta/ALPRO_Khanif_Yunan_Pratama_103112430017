package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan data
	var (
		jamKerja   float64
		upah       float64
		jamNormal  float64 = 40
		upahLembur float64 = 1.5
		jamLembur  float64
		totalGaji  float64
	)

	// Inputan untuk mengisi total jam kerja setiap minggu dan upah per jam
	fmt.Print("Masukan total jam kerja/minggu: ")
	fmt.Scan(&jamKerja)

	fmt.Print("Masukan upah/jam: ")
	fmt.Scan(&upah)

	// Percabangan untuk melakukan sebuah tindakan, ketika jamkerja lebih dari jam normal, maka akan dianggap lembur dan mendapat tambahan gaji.
	if jamKerja > jamNormal {
		jamLembur = jamKerja - jamNormal

		// Rumus yang digunakan untuk menghitung total gaji beserta bonus ketika lembur.
		totalGaji = (jamNormal * upah) + (jamLembur * upahLembur * upah)
		fmt.Printf("Gaji mingguan kamu adalah Rp%.2f\n", totalGaji)
		fmt.Printf("Dalam 1 minggu, kamu memiliki total lembur %.2f jam.\n", jamLembur)
	} else {
		// Ketika tidak lembur, maka akan tampil gaji per minggunya.
		totalGaji = jamKerja * upah

		fmt.Printf("Gaji mingguan kamu adalah Rp%.2f\n", totalGaji)
		fmt.Println("Kamu tidak memiliki jam lembur.")
	}

	// Digunakan untuk menghitung total gaji selama 1 bulan, dikali 4 karena dalam 1 bulan terdapat 4 minggu.
	gajiBulanan := 4 * totalGaji
	fmt.Printf("Total gaji bulanan kamu adalah Rp%.2f\n", gajiBulanan)
}
