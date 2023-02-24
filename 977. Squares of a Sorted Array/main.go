package main

import "fmt"

// [-4,-1,0,3,10]
func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	left, right := 0, len(A)-1
	for i := len(A) - 1; i >= 0; i-- {
		if abs(A[left]) < abs(A[right]) {
			result[i] = A[right] * A[right]
			right--
		} else {
			result[i] = A[left] * A[left]
			left++
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	a := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(a))
}
