package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

// T1
func main() {
	channel := make(chan int)

	poolSize := 100000

	for i := 0; i < poolSize; i++ {
		go worker(i, channel) // 2k cada thread
	}

	for i := 0; i < 10000000; i++ {
		channel <- i
	}
}
