package main

import "fmt"

func main() {

	var tanganSatu, tanganDua, tanganTiga string

	fmt.Scan(&tanganSatu, &tanganDua, &tanganTiga)

	if tanganSatu == tanganDua && tanganDua == tanganTiga && tanganSatu == tanganTiga {
		fmt.Println("imbang")
	} else if tanganSatu != tanganDua && tanganDua == tanganTiga {
		fmt.Print("pemain 1 pemenang")
	} else if tanganDua != tanganSatu && tanganSatu == tanganTiga {
		fmt.Print("pemain 2 pemenang")
	} else if tanganTiga != tanganSatu && tanganSatu == tanganDua {
		fmt.Print("pemain 3 pemenang")
	} else {
		fmt.Print("masukan belum benar")
	}

}
