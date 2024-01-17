package main

import (
	"fmt"
	"math/big"
)


var a = int64(400000000 * 1000000)
var b = int64(500000000 * 1000000)


type BigInt big.Int

func MakeBigInt(num int64) *BigInt {
	return (*BigInt)(big.NewInt(num))
}

func (i1 *BigInt) Add(i2 *BigInt) *BigInt {
	c := new(big.Int).Add((*big.Int)(i1), (*big.Int)(i2))
	return (*BigInt)(c)
}

func (i1 *BigInt) Sub(i2 *BigInt) *BigInt {
	c := new(big.Int).Sub((*big.Int)(i1), (*big.Int)(i2))
	return (*BigInt)(c)
}

func (i1 *BigInt) Mul(i2 *BigInt) *BigInt {
	c := new(big.Int).Mul((*big.Int)(i1), (*big.Int)(i2))
	return (*BigInt)(c)
}

func (i1 *BigInt) Div(i2 *BigInt) *BigInt {
	c := new(big.Int).Div((*big.Int)(i1), (*big.Int)(i2))
	return (*BigInt)(c)
}

func (i1 *BigInt) String() string {
	return (*big.Int)(i1).String()
}

func main() {
	n1 := MakeBigInt(a)
	n2 := MakeBigInt(b)
	fmt.Println("n1 = ", n1)
	fmt.Println("n2 = ", n2)

	add := n1.Add(n2)
	sub := n1.Sub(n2)
	mul := n1.Mul(n2)
	div := n1.Div(n2)
	fmt.Println("add = ", add)
	fmt.Println("sub = ", sub)
	fmt.Println("mul = ", mul)
	fmt.Println("div = ", div)
}