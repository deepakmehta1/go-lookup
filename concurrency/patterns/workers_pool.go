/*
Scenario: Web Server with Worker Pool
Description: A web server routes incoming HTTP requests to a pool of workers (consumers) for processing.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Start 3 consumers (workers)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Producer: Job generator
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
