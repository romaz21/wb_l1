package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, dataChannel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				fmt.Printf("Worker %d exiting...\n", id)
				return
			}
			fmt.Printf("Worker %d received: %d\n", id, data)
		}
	}
}

func main() {
	const numWorkers = 3

	dataChannel := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		close(dataChannel)
	}()

	for i := 1; ; i++ {
		dataChannel <- i
	}

	wg.Wait()
}

