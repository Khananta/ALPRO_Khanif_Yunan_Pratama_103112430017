package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		x1, x2, x3             float64
		y1, y2, y3             float64
		terpanjang             float64
		sisiAB, sisiBC, sisiCA float64
	)

	// Membuat inputan untuk memasukan sumbu x1,2,3 dan y1,2,3
	fmt.Println("Masukan titik X1 dan Y1: ")
	fmt.Scan(&x1, &y1)

	fmt.Println("Masukan titik X2 dan Y2: ")
	fmt.Scan(&x2, &y2)

	fmt.Println("Masukan titik X3 dan Y3: ")
	fmt.Scan(&x3, &y3)

	// Melakukan pengolahan yang dimana math sqrt untuk akar, dan math pow itu sendiri untuk pangkat 2
	sisiAB = math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
	sisiBC = math.Sqrt(math.Pow((x3-x2), 2) + math.Pow((y3-y2), 2))
	sisiCA = math.Sqrt(math.Pow((x1-x3), 2) + math.Pow((y1-y3), 2))

	// Percabangan untuk menentukan sisi terpanjang
	if sisiAB > terpanjang {
		terpanjang = sisiAB
	}
	if sisiBC > terpanjang {
		terpanjang = sisiBC
	}

	if sisiCA > terpanjang {
		terpanjang = sisiCA
	}

	fmt.Printf("Sisi terpanjang: %.2f", terpanjang)
}
