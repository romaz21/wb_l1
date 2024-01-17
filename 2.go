package main

import (
	"fmt"
	"sync"
)

func calculateSquare(wg *sync.WaitGroup, num int) {
	defer wg.Done()

	result := num * num
	fmt.Printf("%d square %d\n", num, result)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(&wg, num)
	}

	wg.Wait()
}

