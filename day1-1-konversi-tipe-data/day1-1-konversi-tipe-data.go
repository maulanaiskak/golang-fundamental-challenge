package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "20"
	b := 25

	fmt.Println("\nKonversi Tipe Data")
	fmt.Println(strings.Repeat("-", 50))

	// print nilai a, b
	fmt.Printf("\nNilai awal A\t: %s dengan tipe data %T \n", a, a)
	fmt.Printf("Nilai awal B\t: %d dengan tipe data %T \n\n", b, b)

	// konversi a menjadi integer dan simpan di variable a_int
	a_int, _ := strconv.Atoi(a)

	// ubah nilai a_int jadi b dan b jadi a_int
	// buat variabel perantara temp untuk tampung nilai a_int
	temp := a_int
	a_int = b
	b = temp

	// print nilai a, b setelah konversi
	fmt.Printf("Nilai akhir A\t: %d dengan tipe data %T \n", a_int, a_int)
	fmt.Printf("Nilai akhir B\t: %d dengan tipe data %T\n\n", b, b)
}
