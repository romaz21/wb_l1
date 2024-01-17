package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped.")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Main goroutine: Canceling worker...")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine: Done.")
}

