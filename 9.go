package main

import (
	"fmt"
	"math/rand"
)

func insertInt(r chan<- int) {
	for {
		rnd := rand.Intn(100)
		r <- rnd
		if rnd == 30 {
			close(r)
			return
		}
	}
}

func checkEven(r <-chan int, e chan<- int) {
	for val := range r {
			e <- val*val
	}
	close(e)
}

func main() {
	r := make(chan int)
	e := make(chan int)

	go insertInt(r)
	go checkEven(r, e)

	for val := range e {
		fmt.Println(val)
	}
}