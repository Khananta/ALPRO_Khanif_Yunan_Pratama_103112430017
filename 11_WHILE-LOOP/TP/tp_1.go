package main

import "fmt"

func main() {

	// Membuat variabel
	var (
		password, inputPassword string
		maxInput                int
	)

	// Isi variabel untuk password dan jumlah maximum memasukan password
	password = "yunan"
	maxInput = 3

	// Perulangan untuk memasukan password, yaitu maksimal 3 kali
	for i := 0; i < maxInput; i++ {

		// Inputan untuk memasukan passoword
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&inputPassword)

		// Pecabangan untuk memeriksa apakah password yang dimasukan benar atau salah.
		if inputPassword == password {
			fmt.Print("Berhasil Masuk")
			break
		} else {
			fmt.Println("Password Salah")
		}

		// Percabangan lagi untuk menampilkan pesan jika upaya login melebihi batas
		if i == maxInput-1 {
			print("Login Ditolak")
		}

	}
}
