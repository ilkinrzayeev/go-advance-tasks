package main

import (
	"fmt"
	"time"
)

func BurstyRateLimiter(requestChan chan bool, resultChan chan int, batchSize int, toAdd int) {
	current := 0
	for {
		ok := <-requestChan
		if !ok {
			return
		}

		for i := 0; i < batchSize; i++ {
			resultChan <- current
			current += toAdd
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	var skipBatches, printBatches, batchSize, toAdd int
	fmt.Scan(&skipBatches)
	fmt.Scan(&printBatches)
	fmt.Scan(&batchSize)
	fmt.Scan(&toAdd)

	requestChan := make(chan bool)
	resultChan := make(chan int)

	go BurstyRateLimiter(requestChan, resultChan, batchSize, toAdd)

	// skip edilməli batch-ləri atla
	for i := 0; i < skipBatches; i++ {
		requestChan <- true
		for j := 0; j < batchSize; j++ {
			<-resultChan
		}
	}

	// çap ediləcək batch-lər
	for i := 0; i < printBatches; i++ {
		requestChan <- true
		for j := 0; j < batchSize; j++ {
			fmt.Println(<-resultChan)
		}
		// yalnız sonuncudan sonra çap etmə
		if i != printBatches-1 {
			fmt.Println("-----")
		}
	}

	close(requestChan)
}
