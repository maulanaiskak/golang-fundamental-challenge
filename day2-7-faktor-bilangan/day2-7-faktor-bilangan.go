package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		bilangan int
		factorStr string
	)

	fmt.Println("\nFaktor Bilangan")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	fmt.Print("Masukkan Bilangan: ")
	_, err := fmt.Scanln(&bilangan)
	fmt.Println()

	
	if err == nil {
		// batas atas looping hanya sampai bilangan/2 + 1 agar looping tidak terlalu banyak
		for i := 1; i <= bilangan/2 + 1; i++ {
			if bilangan % i == 0 {
				factorStr += strconv.Itoa(i) + ", "
			}
		}

		// factorStr ditambahkan nilai bilangan itu sendiri sebagai faktor terakhir, karena pengecekan hanya sampai bilangan/2 + 1
		factorStr += strconv.Itoa(bilangan)
		fmt.Printf("Faktor dari %d: %s\n", bilangan, factorStr)
	} else {
		fmt.Println("Angka harus bulat positif dan lebih dari 1")
	}
	fmt.Println()
}
