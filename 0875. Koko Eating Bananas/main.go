package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, max(piles)

	for left < right {
		mid := (left + right) / 2
		hours := 0

		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(mid)))
		}

		if hours <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func max(arr []int) int {
	max := arr[0]

	for _, val := range arr {
		if max < val {
			max = val
		}
	}
	return max
}
func main() {
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
