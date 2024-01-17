package main

import "fmt"

var slice = []int{0, 1, 2, 3, 4, 5, 6, 7}

func main() {
	fmt.Println(slice)
	copy(slice[4:], slice[5:])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}