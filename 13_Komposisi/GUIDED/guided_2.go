package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Scan(&a, &b, &c)

	max := a
	min := a

	// Percabangan untuk menemukan mana yang lebih besar dan mana yang lebih kecil
	if b > max {
		max = b
	} else if b < min {
		min = b
	}

	if c > max {
		max = c
	} else if c < min {
		min = c
	}

	fmt.Println("Terbesar", max)
	fmt.Println("Terkecil", min)
}
