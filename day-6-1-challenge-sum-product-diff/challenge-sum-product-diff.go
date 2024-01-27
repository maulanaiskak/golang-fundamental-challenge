package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Sum, Product, Difference from 2 Numbers")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println()

	var number string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan 2 angka (pisahkan dengan koma): ")
	number, _ = reader.ReadString('\n')
	number = strings.TrimSpace(number)
	number = strings.ReplaceAll(number, " ", "")

	numberSlice := strings.Split(number, ",")
	if len(numberSlice) == 2 {
		num1, err1 := strconv.Atoi(numberSlice[0])
		num2, err2 := strconv.Atoi(numberSlice[1])

		if err1 == nil && err2 == nil {
			sum, product, diff := operation(num1, num2)
			if num2 < 0 {
				fmt.Printf("%d + (%d) = %d\n", num1, num2, sum)
				fmt.Printf("%d * (%d) = %d\n", num1, num2, product)
				fmt.Printf("%d - (%d) = %d\n", num1, num2, diff)
			} else {
				fmt.Printf("%d + %d = %d\n", num1, num2, sum)
				fmt.Printf("%d * %d = %d\n", num1, num2, product)
				fmt.Printf("%d - %d = %d\n", num1, num2, diff)
			}
		} else {
			fmt.Println("Angka yang anda masukkan salah")
		}
	} else {
		fmt.Println("Angka yang anda masukkan salah")
	}
}

func operation(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	product := num1 * num2
	diff := num1 - num2

	return sum, product, diff
}
