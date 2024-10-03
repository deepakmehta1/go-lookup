/*
Scenario: Task Queue System
Description: Multiple clients (producers) submitting tasks to a queue which are processed by a pool of workers (consumers).
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer function to generate tasks
func producer(id int, tasks chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		task := rand.Intn(100)
		fmt.Printf("Producer %d submitted task %d\n", id, task)
		tasks <- task
		time.Sleep(time.Millisecond * 500) // Simulating some delay
	}
}

// Consumer function to process tasks
func consumer(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Consumer %d processing task %d\n", id, task)
		time.Sleep(time.Millisecond * 500) // Simulating some processing time
	}
}

func main() {
	tasks := make(chan int, 20) // Buffered channel to hold tasks
	var wg sync.WaitGroup

	// Starting producers
	numProducers := 3
	for p := 1; p <= numProducers; p++ {
		wg.Add(1)
		go producer(p, tasks, &wg)
	}

	// Wait for all producers to finish
	wg.Wait()

	// Close tasks channel to signal consumers all tasks are done
	close(tasks)

	// Starting consumers
	numConsumers := 3
	// Wait for all consumers to finish
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(i, tasks, &wg)
	}

	wg.Wait()
}
