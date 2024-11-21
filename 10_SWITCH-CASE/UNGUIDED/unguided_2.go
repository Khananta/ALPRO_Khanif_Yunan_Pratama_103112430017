package main

import "fmt"

func main() {

	// Membuat variable
	var (
		kendaraan     string
		durasi, biaya int
	)

	// Meminta untuk memasukan jenis kendaraan dan durasi
	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)

	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	// Switch case untuk menentukan tarif sesuai dengan jenis kendaraan
	switch kendaraan {
	case "Motor", "motor", "MOTOR":
		biaya = 2000
	case "Mobil", "mobil", "MOBIL":
		biaya = 5000
	case "Truk", "truk", "TRUK":
		biaya = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
	}

	biaya = durasi * biaya

	// Menampilkan hasil biaya parkir yang harus dikeluarkan setiap kendaraan
	fmt.Println("Rp", biaya)

}
