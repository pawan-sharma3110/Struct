package main

import (
	"fmt"
)

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func main() {

	rect := Rectangle{Length: 5, Width: 3}

	area := rect.Area()
	perimeter := rect.Perimeter()

	fmt.Println("Rectangle Area:", area)
	fmt.Println("Rectangle Perimeter:", perimeter)
}
