package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6, 7}, 8))
}
