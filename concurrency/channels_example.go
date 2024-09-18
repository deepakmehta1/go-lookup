package main

import (
	"fmt"  // Importing fmt for printing output
	"sync" // Importing sync to use WaitGroup for goroutine synchronization
)

func main() {
	var wg sync.WaitGroup   // Declare a WaitGroup to synchronize the goroutines
	ch := make(chan string) // Create a channel of string type to send and receive data between goroutines

	wg.Add(1) // Increment the WaitGroup counter by 1 to indicate that we're waiting for one goroutine

	// First goroutine that sends a message to the channel
	go func() {
		ch <- "the message" // Send the string "the message" into the channel
	}()

	// Second goroutine that receives the message from the channel and prints it
	go func() {
		fmt.Println(<-ch) // Receive the message from the channel and print it
		wg.Done()         // Signal that this goroutine is done by decrementing the WaitGroup counter
	}()

	wg.Wait() // Block the main goroutine until the WaitGroup counter goes to 0, ensuring all goroutines finish
}
