package main

import (
	"fmt"
)

func main() {

	// Membuat wadah untuk menyimpan data suhu dengan tipe data float 32
	var celcius float32

	// Inputan yang digunakan untuk memasukan suhu dalam bentuk celcius.
	fmt.Print("Masukan suhu dalam celcius: ")
	fmt.Scan(&celcius)

	//Operasi untuk mengkonfersi dari celcius ke fahrenheit, reamur dan kelvin.
	fahrenheit := (celcius * 9 / 5) + 32
	reamur := celcius * 4 / 5
	kelvin := celcius + 273

	//Menampilkan hasil dari operasi yang telah dilakukan.
	fmt.Println("Temperatur Celcius: ", celcius)
	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Kelvin: ", kelvin)

}
