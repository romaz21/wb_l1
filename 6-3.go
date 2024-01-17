package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(done *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for !*done {
		fmt.Println("Working...")
		time.Sleep(time.Second)
	}
	fmt.Println("Worker stopped.")
}

func main() {
	var done bool
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(&done, &wg)

	time.Sleep(5 * time.Second)
	done = true
	wg.Wait()

	fmt.Println("Main goroutine: Done.")
}

