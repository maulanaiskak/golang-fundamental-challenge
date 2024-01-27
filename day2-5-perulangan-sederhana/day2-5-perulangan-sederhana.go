package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int

	fmt.Println("\nPerulangan Sederhana")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	fmt.Print("Banyak Perulangan: ")
	_, err := fmt.Scanln(&n)
	fmt.Println()

	if err == nil && n > 0 {
		for i := n; i > 0; i-- {
			fmt.Printf("%d - I will become a go developer", i)
			fmt.Println()
		}
	} else {
		fmt.Println("Banyak pengulangan harus bilangan bulat positif")
	}
	fmt.Println()
}
