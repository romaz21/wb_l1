package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int, duration time.Duration) {
	defer close(ch)
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

func receiver(ch <-chan int, done chan<- struct{}) {
	defer close(done)
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

func main() {
	const N = 5

	dataChannel := make(chan int)
	done := make(chan struct{})

	go sender(dataChannel, time.Second)

	go receiver(dataChannel, done)

	time.Sleep(N * time.Second)

	close(dataChannel)
	<-done
	fmt.Println("End")
}

