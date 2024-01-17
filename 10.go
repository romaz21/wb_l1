package main

import (
	"fmt"
	"math"
	"sort"
)

func groupTemperatures(temperatures []float64, step float64) map[float64][]float64 {
	groups := make(map[float64][]float64)

	for _, temp := range temperatures {
		key := math.Floor(temp/step) * step
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	groups := groupTemperatures(temperatures, step)

	var keys []float64
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Float64s(keys)

	for _, key := range keys {
		fmt.Println(key, groups[key])
	}
}

