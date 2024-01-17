package main

import "fmt"

func createSet(sequence []string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, element := range sequence {
		set[element] = struct{}{}
	}

	return set
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	mySet := createSet(sequence)
	fmt.Println("Множество:", mySet)
}

