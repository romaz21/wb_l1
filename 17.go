package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{11, 12, 22, 25, 34, 64, 90}

	target := 22

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
	if index < len(arr) && arr[index] == target {
		fmt.Println(target, index)
	} else {
		fmt.Println(false)
	}
}

