package main

import (
	"fmt"
)

func main() {

	// Membuat variabel
	var jmlPenumpang int

	// Inputan untuk memasukan jumlah penumpang
	fmt.Print("Masukan Jumlah Penumpang: ")
	fmt.Scan(&jmlPenumpang)

	// Percabangan untuk menentukan jumlah motor
	if jmlPenumpang%2 == 0 {
		jmlPenumpang = jmlPenumpang / 2
	}

	// Output yang ditampilkan ketika percabangan tidak terpenuhi
	fmt.Print(jmlPenumpang)

}
