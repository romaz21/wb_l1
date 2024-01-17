package main

import "fmt"

var sls = []string{
	"abcd",
	"abCdefAaf",
	"aabcd",
}

func check(s string) bool {
	runes := []rune(s)
	m := make(map[rune]int)
	for _, rune := range runes {
		if _, ok := m[rune]; ok {
			return false
		} else {
			m[rune] = 1
		}
	}
	return true
}

func main() {
	for _, s := range sls {
		fmt.Println(s, check(s))
	}
}