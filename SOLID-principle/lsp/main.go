package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side // Correct behavior for a square
}

func PrintArea(shape Shape) {
	fmt.Println("Area:", shape.Area())
}

func main() {
	rectangle := &Rectangle{Width: 5, Height: 10}
	square := &Square{Side: 4}

	PrintArea(rectangle) // Outputs: 50
	PrintArea(square)    // Outputs: 16
}
