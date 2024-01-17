package main

import "fmt"

func swapWithoutTemp(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	num1, num2 := 5, 10

	fmt.Println(num1, num2)

	num1, num2 = swapWithoutTemp(num1, num2)

	fmt.Println(num1, num2)
}

