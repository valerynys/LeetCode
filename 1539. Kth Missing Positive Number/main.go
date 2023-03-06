package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	last := 0
	for _, num := range arr {
		diff := num - last - 1
		if diff >= k {
			return last + k
		}
		k -= diff
		last = num
	}
	return last + k
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 5, 6, 7, 8, 9, 17, 23, 42}, 10))
}
