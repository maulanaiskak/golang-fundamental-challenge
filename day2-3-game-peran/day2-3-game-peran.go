package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		nama, peran, superhero, monster string
	)

	fmt.Println("\nGame Nama dan Peran")
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()

	fmt.Printf("Masukkan nama\t: ")
	fmt.Scanln(&nama)

	fmt.Printf("Masukkan peran\t: ")
	fmt.Scanln(&peran)

	// jadikan peran lowercase karena tidak perlu case sensitive
	peran = strings.ToLower(peran)
	superhero = "superhero"
	monster = "monster"
	fmt.Println()

	if nama == "" {
		if peran != "" {
			fmt.Println("Nama dan peran harus diisi")
		} else {
			fmt.Println("Nama harus diisi")
		}
	} else {
		if peran == "" {
			fmt.Println("Peran harus diisi")
		} else if peran == superhero {
			fmt.Println("Selamat Datang Superhero Saitama, Kalahkan Semua Monster di Muka Bumi")
		} else if peran == monster {
			fmt.Println("Selamat Datang Monster Saitama, Hancurkan Semua Superhero yang Ada")
		} else {
			fmt.Println("Selamat Datang Saitama, Pilih Peranmu Untuk Melanjutkan Game Ini")
		}
	}
	fmt.Println()
}
