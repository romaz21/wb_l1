package main

import (
	"fmt"
	"strings"
)

func ReverseWords(input string) string {
	words := strings.Fields(input)
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	return strings.Join(reversedWords, " ")
}

func main() {
	inputString := "snow dog sun — sun dog snow"
	reversedString := ReverseWords(inputString)

	fmt.Println(inputString)
	fmt.Println(reversedString)
}

