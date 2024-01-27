package main

import (
	"challenge-method-persegi-lingkaran/calc"
	"fmt"
)

func main() {
	rectangle := calc.Rectangle{Width: 12.5, Height: 9.8}
	fmt.Printf("Luas dari persegi: %.3f\n\n", rectangle.Area())

	circle := calc.Circle{Radius: 12.5}
	fmt.Printf("Luas dari lingkaran: %.3f\n\n", circle.Area())
}
