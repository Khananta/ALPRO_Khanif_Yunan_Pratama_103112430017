package main

import "fmt"

//Function yang digunakan untuk menghitung apakah tahun yang kita masukan merupakan tahun kabisat.
func kabisat(tahun int) bool {
	// Ketika tahun yang kita masukan memenuhi syarat yaitu habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100 maka output yang ditampilkan adalah true.
	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		return true
	}
	return false //Ketika tidak sesuai dengan aturan yang telah dibuat maka output akan bernilai false.
}

func main() {

	var tahun int //Membuat wadah untuk menampung tahun yang dimasukan.

	//Inputan untuk memasukan data tahun.
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	//Memanggil function yang telah dibuat dan menampilkan.
	fmt.Print("Kabisat: ", kabisat(tahun))

}
