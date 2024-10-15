package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Fungsi untuk memilih angka acak antara 1 hingga 100
	rand.Seed(time.Now().UnixNano())
	angkaRahasia := rand.Intn(100) + 1

	// Variabel untuk menyimpan data tebakan
	var tebakan int

	// Looping yang digunakan untuk percobaan, yaitu maximal 5 kali percobaan
	for percobaan := 1; percobaan <= 5; percobaan++ {
		fmt.Printf("(Kesempatan %d) Masukkan angka kamu: ", percobaan)
		fmt.Scan(&tebakan)

		// Percabangan untuk menunjukan angka yang kita masukan apakah terlalu besar atau kecil
		if tebakan < angkaRahasia {
			fmt.Println("Terlalu kecil!")
			fmt.Print("--------------\n")
		} else if tebakan > angkaRahasia {
			fmt.Println("Terlalu besar!")
			fmt.Print("--------------\n")
		} else {
			fmt.Printf("Selamat! Tebakan kamu benar. Angkanya adalah %d.\n", angkaRahasia)
			return
		}
	}

	// Output yang akan tampil ketika kesempatan menjawab sudah habis.
	fmt.Printf("Maaf, kesempatan habis. Angka yang benar adalah %d.\n", angkaRahasia)
}
