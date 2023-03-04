package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 || k > n {
		return []int{}
	}
	res := make([]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		res[i] = max(nums[i : i+k])
	}
	return res
}

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

func main() {
	nums := []int{2, -1, -4, 6, 8, 9, 11}
	k := 2
	fmt.Println(maxSlidingWindow(nums, k))
}
