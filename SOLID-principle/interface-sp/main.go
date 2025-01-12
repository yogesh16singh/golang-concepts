package main

import "fmt"

// Smaller interfaces focused on specific roles
type Coder interface {
	Code()
}

type Tester interface {
	Test()
}

type Designer interface {
	Design()
}

type Developer struct{}

func (d *Developer) Code() {
	fmt.Println("Writing code")
}

func (d *Developer) Test() {
	fmt.Println("Testing code")
}

type GraphicDesigner struct{}

func (gd *GraphicDesigner) Design() {
	fmt.Println("Creating designs")
}

func main() {
	// A developer can code and test
	dev := &Developer{}
	dev.Code()
	dev.Test()

	// A designer only designs
	gd := &GraphicDesigner{}
	gd.Design()
}
