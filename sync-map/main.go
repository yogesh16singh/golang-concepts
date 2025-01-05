package main

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var cache = &sync.Map{}

func calculateFibonacci(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	a := big.NewInt(0)
	b := big.NewInt(1)
	var result *big.Int

	for i := 2; i <= n; i++ {
		result = new(big.Int).Set(a)
		result.Add(result, b)
		a.Set(b)
		b.Set(result)
	}

	return result
}

func fibonacciHandler(c *gin.Context) {
	n := c.DefaultQuery("n", "0")
	nInt, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the result is already in the cache
	if cacheResult, ok := cache.Load(nInt); ok {
		fmt.Println("Cache hit!")
		c.JSON(http.StatusOK, gin.H{"result": cacheResult})
		return
	}

	fmt.Println("Cache miss!")

	// Calculate Fibonacci if not in cache
	result := calculateFibonacci(nInt)

	// Store the result in the cache with a time-to-live (TTL)
	cache.Store(nInt, result)
	go func() {
		// Remove the entry from cache after 5 minutes (300 seconds)
		time.Sleep(5 * time.Minute)
		cache.Delete(nInt)
	}()

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func main() {
	r := gin.Default()

	r.GET("/fibonacci", fibonacciHandler)

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	r.Run(port)
}
