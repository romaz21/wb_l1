package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Before:", arr)
	sort.Sort(IntSlice(arr))
	fmt.Println("After:", arr)
}

