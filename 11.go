package main

import "fmt"

func intersection(set1, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for key := range set1 {
		if _, exists := set2[key]; exists {
			result[key] = struct{}{}
		}
	}

	return result
}

func main() {
	set1 := map[int]struct{}{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
		4: struct{}{},
	}

	set2 := map[int]struct{}{
		3: struct{}{},
		4: struct{}{},
		5: struct{}{},
		6: struct{}{},
	}

	result := intersection(set1, set2)
	fmt.Println(result)
}

