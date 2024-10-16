package main

import (
	"fmt"
)

func main() {

	var (
		a, b int
	)

	fmt.Print("Masukan a: ")
	fmt.Scan(&a)

	fmt.Print("Masukan b: ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		fmt.Print(i)
	}

}
