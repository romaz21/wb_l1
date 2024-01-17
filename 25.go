package main

import (
	"fmt"
	"time"
)

func Sleep(milliseconds time.Duration) {
	timer := time.NewTicker(milliseconds * time.Millisecond)

	<-timer.C

	timer.Stop()
}

func main() {
	fmt.Println("Begin")
	Sleep(3000)
	fmt.Println("End")
}
