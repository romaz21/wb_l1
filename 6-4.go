package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker() {
	for {
		fmt.Println("Working...")
		time.Sleep(time.Second)
	}
}

func main() {
	go worker()

	time.Sleep(2 * time.Second)

	fmt.Println("Main goroutine: Stopping worker...")
	runtime.Goexit()

	fmt.Println("Main goroutine: Done.")
}

