package main

import "fmt"

func main() {

	var jumlahDaun, lebar, lebarTerbesar int

	fmt.Scan(&jumlahDaun)

	lebarTerbesar = 0

	for i := 1; i <= jumlahDaun; i++ {
		fmt.Scan(&lebar)

		if lebar > lebarTerbesar {
			lebarTerbesar = lebar
		}
	}

	fmt.Print(lebarTerbesar)
}
