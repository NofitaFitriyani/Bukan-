package main

import (
	"fmt"
)

func main() {
	var angka int
	total := 0
	fmt.Println("Masukkan angka positif untuk dijumlahkan. Masukan angka negatif untuk keluar.")
	for {
		fmt.Print("Masukkan angka: ")
		fmt.Scanln(&angka)
		if angka < 0 {
			break
		}
		total += angka
	}
	fmt.Println("Total jumlah dari angka yang dimasukan adalah: %d\n", total)
}
