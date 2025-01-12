package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a unit of work to be done.
type Task struct {
	ID       int
	Workload int
}

// Worker processes tasks from the task channel.
func Worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task %d with workload %d\n", id, task.ID, task.Workload)
		time.Sleep(time.Duration(task.Workload) * time.Millisecond * 10) // Simulate work
		fmt.Printf("Worker %d finished task %d\n", id, task.ID)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Configuration
	numWorkers := 3
	numTasks := 10

	// Channels
	taskChannel := make(chan Task)

	// WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, taskChannel, &wg)
	}

	// Generate and send tasks to the task channel
	go func() {
		for i := 1; i <= numTasks; i++ {
			task := Task{ID: i, Workload: rand.Intn(500) + 100} // Random workload
			taskChannel <- task
		}
		close(taskChannel) // Close the task channel to signal workers to stop
	}()

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All tasks processed!")
}
