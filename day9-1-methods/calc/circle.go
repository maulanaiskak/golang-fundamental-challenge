package calc

import (
	"math"
)

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	area := math.Pi * math.Pow(c.Radius, 2)
	// area := 3.14 * c.Radius * c.Radius
	return area
}
