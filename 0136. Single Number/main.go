package main

import "fmt"

func singleNumber(nums []int) int {
	var sol int
	for i := 0; i < len(nums); i++ {
		sol = sol ^ nums[i]
		fmt.Println(sol)
	}

	return sol
}

func main() {
	arr := []int{1, 2, 4, 2, 4}
	fmt.Println(singleNumber(arr))
}
