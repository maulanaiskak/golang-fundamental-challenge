package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Even and Odd Numbers")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println()

	var numbers []int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan kumpulan angka (pisahkan dengan koma): ")
	numbersStr, _ := reader.ReadString('\n')
	numbersStr = strings.TrimSpace(numbersStr)
	numbersStr = strings.ReplaceAll(numbersStr, " ", "")
	numbersSlice := strings.Split(numbersStr, ",")

	eror := false
	for _, numberStr := range numbersSlice {
		num, err := strconv.Atoi(numberStr)
		if err == nil {
			numbers = append(numbers, num)
			eror = false
		} else {
			eror = true
			break
		}
	}

	if !eror {
		even, odd := filter(numbers, isEven)
		fmt.Printf("\nEven numbers\t: %v\n", even)
		fmt.Printf("\nOdd numbers\t: %v\n", odd)
	} else {
		fmt.Println("\nAngka yang anda masukkan salah")
	}
	fmt.Println()
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func filter(numbers []int, callback func(int) bool) (even, odd []int) {
	for _, number := range numbers {
		if callback(number) {
			even = append(even, number)
		} else {
			odd = append(odd, number)
		}
	}
	return even, odd
}
