package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i1, i2 := m-1, n-1
	for end := len(nums1) - 1; end >= 0 && i2 >= 0; end-- {
		if i1 >= 0 && nums1[i1] > nums2[i2] {
			nums1[end] = nums1[i1]
			i1--
		} else {
			nums1[end] = nums2[i2]
			i2--
		}
	}
	return nums1
}

func main() {
	arr1 := []int{1, 2, 6, 8, 21, 4, 3, 7, 8, 4}
	arr2 := []int{11, 2, 6, 28, 215, 4, 3, 17, 68, 40}
	fmt.Println(merge(arr1, 5, arr2, 5))
}
