package main

import (
	"fmt"
	"time"
)

func worker(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Worker stopped.")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	stopCh := make(chan struct{})

	go worker(stopCh)

	time.Sleep(5 * time.Second)
	close(stopCh)

	fmt.Println("Main goroutine: Waiting for worker to stop...")
	time.Sleep(2 * time.Second)
	fmt.Println("Main goroutine: Done.")
}

