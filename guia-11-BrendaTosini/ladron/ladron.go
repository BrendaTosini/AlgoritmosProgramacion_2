package main

import (
	"fmt"
	"math"
)

func Rob(houses []int) int {
	return rob(houses, 0)
}

func rob(houses []int, current int) int {
	if current >= len(houses) {
		return 0
	}

	return max(houses[current]+rob(houses, current+2), rob(houses, current+1))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	fmt.Println(Rob([]int{2, 10, 3, 6, 8, 1, 7}))
}
