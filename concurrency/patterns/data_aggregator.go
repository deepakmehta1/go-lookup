/*
Scenario: Data Aggregator
Description: Multiple sensors (producers) send data to a central server (consumer) for aggregation.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sensor(id int, dataCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		data := rand.Intn(100)
		fmt.Printf("Sensor %d generated data %d\n", id, data)
		dataCh <- data
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	dataCh := make(chan int, 20)
	var wg sync.WaitGroup

	// Start 3 producers (sensors)
	for s := 1; s <= 3; s++ {
		wg.Add(1)
		go sensor(s, dataCh, &wg)
	}

	// Consumer: Data aggregator
	go func() {
		for data := range dataCh {
			fmt.Println("Aggregating data: ", data)
		}
	}()

	wg.Wait()
	close(dataCh)
	time.Sleep(time.Second * 2) // Waiting for aggregator to finish
}
