package main

import "fmt"

func main() {

	var angka int
	var kebenaran bool

	kebenaran = true

	fmt.Scan(&angka)

	if angka <= 1 {
		kebenaran = false
	} else {
		for i := 2; i < angka; i++ {
			if angka%i == 0 {
				kebenaran = false
				break
			}
		}
	}

	fmt.Println(kebenaran)

}
