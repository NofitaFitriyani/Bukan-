package main
import (
	"fmt"
)
func main() {
	maxNumber := 10
	number := 1

	for finished := false; !finished; {
		fmt.Println("Angka: ", number)

		if number >= maxNumber {
			finished = true
		}
		number ++
	}