package main

import (
	"fmt"
)

func SetBit(num *int64, i uint) {
	*num |= 1 << i
}

func ClearBit(num *int64, i uint) {
	*num &^= 1 << i
}

func main() {
	var num int64 = 42

	SetBit(&num, 3)
	fmt.Printf("After setting bit 3: %d\n", num)

	ClearBit(&num, 5)
	fmt.Printf("After clearing bit 5: %d\n", num)
}

