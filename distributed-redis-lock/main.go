package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func acquireLock(client *redis.Client, lockKey string, timeout time.Duration) bool {

	// Try to acquire the lock with SETNX command (SET if Not Exists)
	lockAcquired, err := client.SetNX(ctx, lockKey, "1", timeout).Result()
	if err != nil {
		fmt.Println("Error acquiring lock:", err)
		return false
	}

	return lockAcquired
}

func releaseLock(client *redis.Client, lockKey string) {
	_, err := client.Del(ctx, lockKey).Result()
	if err != nil {
		fmt.Println("Error releasing lock:", err)
	}
}

func main() {
	// Create a Redis client.
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "redis://def", // Redis server address
	// 	Password: "",                                                                                                           // No password by default
	// 	DB:       0,                                                                                                            // Default DB
	// })
	opt, _ := redis.ParseURL("redis url")
	client := redis.NewClient(opt)

	client.Set(ctx, "foo", "bar", 0)
	// val := client.Get(ctx, "foo").Val()

	defer func() {
		if err := client.Close(); err != nil {
			fmt.Println("Error closing Redis client:", err)
		}
	}()

	// Define the lock key and lock timeout
	lockKey := "my_lock"
	lockTimeout := 20 * time.Second

	// Acquire the lock
	if acquireLock(client, lockKey, lockTimeout) {
		fmt.Println("Lock acquired successfully!")
		// Simulate some work with the lock
		time.Sleep(20 * time.Second)
		fmt.Println("Work done!")

		// Release the lock
		releaseLock(client, lockKey)
		fmt.Println("Lock released.")
	} else {
		fmt.Println("Failed to acquire lock. Resource is already locked.")
	}
}
