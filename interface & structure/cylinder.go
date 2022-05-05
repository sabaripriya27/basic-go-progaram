package main

import "fmt"

// Creating an interface
type cylinder interface {

	// Methods
	surfaceArea() float64
	Volume() float64
	lateralsurface() float64
	ToporBotsurface() float64
}

type myvalue struct {
	radius float64
	height float64
}

// Implementing methods of
// the tank interface
func (m myvalue) surfaceArea() float64 {

	return (2 * 3.14 * m.radius * m.radius) + (2*3.14*m.radius + m.height)
}

func (m myvalue) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}
func (m myvalue) lateralsurface() float64 {

	return 2*3.14*m.radius + m.height
}
func (m myvalue) ToporBotsurface() float64 {

	return 3.14 * m.radius * m.radius
}

// Main Method
func main() {

	// Accessing elements of
	// the tank interface
	var c cylinder
	c = myvalue{10, 14}
	fmt.Println("surface Area of cylinder :", c.surfaceArea())
	fmt.Println("Volume of cylinder:", c.Volume())
	fmt.Println("lateral surface area of cylinder:", c.lateralsurface())
	fmt.Println("Top or Bot surface area of cylinder:", c.ToporBotsurface())

}
