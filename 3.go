package main

import (
	"fmt"
	"sync"
)

func calculateSquare(num int, resultChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	square := num * num
	resultChannel <- square
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	resultChannel := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, resultChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	sum := 0
	for square := range resultChannel {
		sum += square
	}

	fmt.Println(sum)
}

