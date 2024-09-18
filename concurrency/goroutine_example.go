package main

import (
	"fmt"  // Importing the fmt package for formatted I/O functions like Println
	"sync" // Importing the sync package to use WaitGroup for synchronizing goroutines
)

func main() {
	var wg sync.WaitGroup // Declare a WaitGroup, used to wait for goroutines to finish

	wg.Add(1) // Increment the WaitGroup counter by 1, indicating we're waiting for 1 goroutine

	// Launch a goroutine to run asynchronously
	go func() {
		fmt.Println("do some async thing") // This code runs concurrently, printing a message
		wg.Done()                          // Decrement the WaitGroup counter by 1, signaling that the goroutine is done
	}()

	wg.Wait() // Block the main goroutine until the WaitGroup counter goes to 0 (all goroutines are done)
}
