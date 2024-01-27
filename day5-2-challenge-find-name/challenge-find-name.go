package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Pencarian Nama")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan kumpulan nama: ")
	namaStr, _ := reader.ReadString('\n')
	namaStr = strings.TrimSpace(namaStr)
	namaSlice := strings.Split(namaStr, " ")

	fmt.Println()
	fmt.Printf("Nama dengan jumlah karakter genap: ")
	for _, nama := range namaSlice {
		if len(nama)%2 == 0 {
			fmt.Printf("%s ", nama)
		}
	}
	fmt.Printf("\n\n")

}
