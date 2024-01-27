package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		sisiBalok, jariJari, alasSegitiga, tinggiSegitiga float32
	)

	fmt.Println("\nMenghitung luas balok, lingkaran, segitiga sama kaki")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Print("\nMasukkan panjang sisi balok (meter)\t: ")
	fmt.Scanln(&sisiBalok)

	fmt.Print("Masukkan jari-jari lingkaran (meter)\t: ")
	fmt.Scanln(&jariJari)

	fmt.Print("Masukkan panjang alas segitiga (meter)\t: ")
	fmt.Scanln(&alasSegitiga)

	fmt.Print("Masukkan tinggi segitiga (meter)\t: ")
	fmt.Scanln(&tinggiSegitiga)

	// hitung luas bangun datar
	luasBalok := sisiBalok * sisiBalok
	luasLingkaran := 3.14 * jariJari * jariJari
	luasSegitiga := alasSegitiga * tinggiSegitiga / 2

	fmt.Printf("\nLuas Balok (m^2)\t\t\t: %.2f \n", luasBalok)
	fmt.Printf("Luas Lingkaran (m^2)\t\t\t: %.2f \n", luasLingkaran)
	fmt.Printf("Luas Segitiga Sama Sisi (m^2)\t\t: %.2f \n\n", luasSegitiga)
}
