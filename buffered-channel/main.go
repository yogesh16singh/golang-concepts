package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string, 3)

	go func() {
		for i := 1; i <= 5; i++ {
			order := fmt.Sprintf("Coffee Order #%d", i)
			orders <- order
			fmt.Println("Order placed:", order)
		}
		close(orders)
	}()

	for order := range orders {
		fmt.Printf("preparing: %s\n", order)
		time.Sleep(2 * time.Second)
		fmt.Printf("prepared: %s\n", order)
	}
}
