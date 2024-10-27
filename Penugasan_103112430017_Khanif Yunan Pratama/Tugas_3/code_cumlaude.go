package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan data dan menyatakan nilai kebenaran
	var (
		cumlaude       bool
		semester, eprt int
	)

	// Input untuk memasukan nilai semester dan eprt
	fmt.Print("Masukan jumlah semester: ")
	fmt.Scanln(&semester)

	fmt.Print("Masukan EpRT: ")
	fmt.Scanln(&eprt)

	// Boolean untuk menyatakan cumlaude jika lulus semester <= 8 dan eprt >= 500
	cumlaude = semester <= 8 && eprt >= 500
	fmt.Println(cumlaude)
}
