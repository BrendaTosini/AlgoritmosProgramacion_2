package main

import (
	"fmt"
	"math"
)

func Rob(houses []int) int {
	n := len(houses)

	table := make([]int, n)
	table[0] = houses[0]
	table[1] = houses[1]

	for x := 2; x < n; x++ {
		table[x] = max(houses[x]+table[x-2], table[x-1])
	}

	return table[n-1]
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	fmt.Println(Rob([]int{2, 10, 3, 6, 8, 1, 7}))
}
