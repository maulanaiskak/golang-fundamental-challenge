package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var kata string
	var palindrome bool

	fmt.Println("\nPalindrome")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&kata)
	_, err := strconv.Atoi(kata)

	fmt.Println()

	kata = strings.ToLower(kata)
	panjangKata := len(kata)

	if err != nil {
		// digunakan indexing, output indexing = ASCII dari huruf
		// batas for loop sampai panjang kata/2 + 1 agar tidak melakukan cek 2x dan lebih efisien

		for i := 0; i < panjangKata/2+1; i++ {
			// jika hasil indexing dari depan == indexing dari belakang --> palindrome

			if kata[i] == kata[panjangKata-i-1] {
				palindrome = true
			} else {
				palindrome = false
			}
		}

		if palindrome {
			fmt.Printf("%s merupakan palindrome\n", kata)
		} else {
			fmt.Printf("%s bukan merupakan palindrome\n", kata)
		}
	} else {
		fmt.Println("Kata harus berupa string")
	}
	fmt.Println()
}
