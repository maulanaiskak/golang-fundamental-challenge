package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		angka, genap        int
		ganjilStr, genapStr string
	)
	jumlahGanjil := 0
	jumlahGenap := 0

	fmt.Println("\nPenjumlahan Ganjil Genap")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println()

	fmt.Print("Masukkan angka: ")
	_, err := fmt.Scanln(&angka)
	fmt.Println()

	// ganjilStr dan genapStr akan menyimpan bilangan ganjil/genap dalam bentuk 1 + 2 + 3 + ....
	if err == nil && angka > 1 {
		for i := 0; i < angka; i++ {
			if i%2 != 0 {
				ganjilStr += strconv.Itoa(i) + " + "
				jumlahGanjil += i
			}
			genap += 2
			genapStr += strconv.Itoa(genap) + " + "
			jumlahGenap += genap
		}

		ganjilStr = strings.TrimSuffix(ganjilStr, " + ")
		genapStr = strings.TrimSuffix(genapStr, " + ")

		fmt.Printf("Penjumlahan angka ganjil dari 0 sampai %d\t: %s = %d\n", angka, ganjilStr, jumlahGanjil)
		fmt.Printf("Penjumlahan %d angka genap pertama\t\t: %s = %d\n\n", angka, genapStr, jumlahGenap)
	} else {
		fmt.Println("Angka harus bulat positif dan lebih dari 1")
	}

	fmt.Println()
}
