package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	start_index := 0
	var result []string

	for i := range nums {
		// Find last index for continuous series
		if i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			continue
		}
		if start_index == i {
			result = append(result, strconv.Itoa(nums[start_index]))
		} else {
			result = append(result, strconv.Itoa(nums[start_index])+"->"+strconv.Itoa(nums[i]))
		}
		start_index = i + 1
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	fmt.Println(summaryRanges(arr))
}
