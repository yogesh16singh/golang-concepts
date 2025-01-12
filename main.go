package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	fmt.Println("Hello, World!")
	testenv := os.Getenv("ENV_TEST")
	fmt.Println(testenv)
}
