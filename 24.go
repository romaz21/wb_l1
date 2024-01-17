package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func Distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	distance := Distance(point1, point2)

	fmt.Println(distance)
}
