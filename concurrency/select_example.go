package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels for integer data
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start a goroutine that sends data to ch1 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work with a delay
		ch1 <- 1                    // Send the integer 1 to ch1
	}()

	// Start another goroutine that sends data to ch2 after 1 second
	go func() {
		time.Sleep(1 * time.Second) // Simulate some work with a delay
		ch2 <- 2                    // Send the integer 2 to ch2
	}()

	// Use select to wait for data from either channel, or handle the default case if no channel is ready
	select {
	case msg1 := <-ch1: // If a message is received from ch1, this case will execute
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2: // If a message is received from ch2, this case will execute
		fmt.Println("Received from ch2:", msg2)
	default:
		// If neither ch1 nor ch2 are ready, the default case executes to prevent blocking (deadlock)
		fmt.Println("No message received yet, avoiding deadlock.")
	}
}
