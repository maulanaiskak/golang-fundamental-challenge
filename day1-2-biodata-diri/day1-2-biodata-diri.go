package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// digunakan bufio untuk input karena kemungkinan akan ada whitespace
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("\nBiodata Diri")
	fmt.Println(strings.Repeat("-", 120))

	fmt.Print("Masukkan Nama Depan \t: ")
	firstName, _ := inputReader.ReadString('\n')

	fmt.Print("Masukkan Nama Belakang \t: ")
	lastName, _ := inputReader.ReadString('\n')

	fmt.Print("Masukkan Alamat \t: ")
	address, _ := inputReader.ReadString('\n')

	fmt.Print("Masukkan Umur \t\t: ")
	age, _ := inputReader.ReadString('\n')

	fmt.Print("Masukkan Skills \t: ")
	skills, _ := inputReader.ReadString('\n')

	//hilangkan enter dari string hasil input bufio
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	address = strings.TrimSpace(address)
	age = strings.TrimSpace(age)
	skills = strings.TrimSpace(skills)

	fullName := firstName + " " + lastName

	fmt.Println("\nBiodata Diri")
	fmt.Println(strings.Repeat("-", 120))

	fmt.Printf("Nama Lengkap \t| Umur \t| Alamat \t\t\t\t| Skills\n")
	fmt.Println(strings.Repeat("-", 120))
	fmt.Printf("%s \t| %v \t| %s \t\t\t| %s\n", fullName, age, address, skills)
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println()

}
