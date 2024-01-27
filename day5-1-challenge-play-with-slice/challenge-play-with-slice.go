package main

import (
	"fmt"
	"strings"
)

func main() {
	var dataKaryawan = [][]string{
		{"A001", "Budi Hartono", "Bandung"},
		{"A002", "Michael Parmon", "Bekasi"},
		{"A003", "Saphira", "Bandung"},
		{"A004", "Santhia", "Jakarta"},
		{"A005", "Ageng", "Depok"},
	}

	fmt.Println("Data Karyawan")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println()

	for _, karyawan := range dataKaryawan {
		fmt.Printf("ID Karyawan\t: %s\n", karyawan[0])
		fmt.Printf("Nama Karyawan\t: %s\n", karyawan[1])
		fmt.Printf("Alamat\t\t: %s\n\n", karyawan[2])
	}
}