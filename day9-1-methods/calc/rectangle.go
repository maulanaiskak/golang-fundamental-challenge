package calc

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	area := r.Width * r.Height
	// fmt.Printf("Luas dari persegi: %.3f\n\n", area)
	return area
}
