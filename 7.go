package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	counts map[string]int
}

func (c *SafeCounter) Increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

func (c *SafeCounter) GetValue(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[key]
}

func main() {
	counter := SafeCounter{counts: make(map[string]int)}

	numWorkers := 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				key := fmt.Sprintf("key%d", workerID)
				counter.Increment(key)
				time.Sleep(time.Millisecond * 100)
			}
		}(i)
	}

	wg.Wait()

	for i := 0; i < numWorkers; i++ {
		key := fmt.Sprintf("key%d", i)
		value := counter.GetValue(key)
		fmt.Printf("Key %s: %d\n", key, value)
	}
}

