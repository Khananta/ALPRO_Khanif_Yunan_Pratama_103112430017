package main

import "fmt"

func main() {

	var kondisiBadan, persediaan, cuaca bool

	fmt.Scan(&kondisiBadan, &persediaan, &cuaca)

	if kondisiBadan == true && persediaan == true && cuaca == true {
		fmt.Println("Berkemah")
	} else {
		fmt.Println("Tidak berkemah")
	}

}
