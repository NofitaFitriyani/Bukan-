package main

import "fmt"

func main() {
	var base, height float64

	// Input alas dan tinggi dari pengguna
	fmt.Print("Masukan nilai alas: ")
	fmt.Scan(&base)

	fmt.Print("Masukan nilai tinggi: ")
	fmt.Scan(&height)

	//Rumus luas segitiga = 1/2 * alas * tinggi
	area := 0.5 * base * height

	// menampilkan hasil perhitungan luas segitiga
	fmt.Printf("Luas segitiga adalah: %.2f\n", area)

}
