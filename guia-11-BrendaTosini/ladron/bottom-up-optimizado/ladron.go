package main

import (
	"fmt"
	"math"
)

func Rob(houses []int) int {
	n := len(houses)

	prevPrev := houses[0]
	prev := houses[1]
	cur := 0

	for x := 2; x < n; x++ {
		cur = max(houses[x]+prevPrev, prev)

		prevPrev = prev
		prev = cur
	}

	return cur
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	fmt.Println(Rob([]int{2, 10, 3, 6, 8, 1, 7}))
}
