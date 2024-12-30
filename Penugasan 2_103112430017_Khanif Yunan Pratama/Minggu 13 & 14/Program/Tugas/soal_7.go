package main

import "fmt"

func main() {
	var angka int
	var max, min, sum, count int
	var selesai bool

	max = -1000000
	min = 1000000

	for {
		fmt.Scan(&angka)
		if angka == 0 {
			if selesai {
				break
			}
			selesai = true
			continue
		}

		selesai = false
		if angka > max {
			max = angka
		}
		if angka < min {
			min = angka
		}
		sum += angka
		count++
	}

	if count > 0 {
		rata := float64(sum) / float64(count)
		fmt.Printf("Tertinggi: %d\n", max)
		fmt.Printf("Terendah: %d\n", min)
		fmt.Printf("Rata-rata: %.3f\n", rata)
	} else {
		fmt.Println("Tidak ada data yang valid.")
	}
}
