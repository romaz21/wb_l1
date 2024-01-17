package main

import (
    "fmt"
    "strings"
)
var justString string

func createHugeString(size int) string {
    return strings.Repeat("huge string", 10)
}

func someFunc() string {
    return createHugeString(1 << 10)[:100]
}

func main() {
    justString = someFunc()
    fmt.Println(justString)
}

