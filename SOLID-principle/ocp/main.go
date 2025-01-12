package main

import "fmt"

// Define a DiscountCalculator interface
type DiscountCalculator interface {
	Calculate(amount float64) float64
}

// Regular customer discount
type RegularDiscount struct{}

func (r *RegularDiscount) Calculate(amount float64) float64 {
	return amount * 0.9 // 10% discount
}

// VIP customer discount
type VIPDiscount struct{}

func (v *VIPDiscount) Calculate(amount float64) float64 {
	return amount * 0.8 // 20% discount
}

// New discount type: Employee
type EmployeeDiscount struct{}

func (e *EmployeeDiscount) Calculate(amount float64) float64 {
	return amount * 0.7 // 30% discount
}

func main() {
	// Use the DiscountCalculator interface
	discounts := []DiscountCalculator{
		&RegularDiscount{},
		&VIPDiscount{},
		&EmployeeDiscount{},
	}

	amount := 100.0
	for _, discount := range discounts {
		fmt.Println("Discounted amount:", discount.Calculate(amount))
	}
}
