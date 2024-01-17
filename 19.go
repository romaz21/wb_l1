package main

import (
	"fmt"
)

func ReverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	inputString := "hello world"
	reversedString := ReverseString(inputString)

	fmt.Println(inputString)
	fmt.Println(reversedString)
}

