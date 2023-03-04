package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	inSection := make([]int, 0)
	fm := make(map[int]int)
	for _, n := range nums1 {
		fm[n]++
	}
	fmt.Println(fm)
	for _, n := range nums2 {
		if v, ok := fm[n]; ok {
			if v != 0 {
				inSection = append(inSection, n)
				fm[n] = v - 1
			}
			fmt.Println(fm)
		}
	}
	fmt.Println(inSection)
	return inSection
}

func main() {
	intersect([]int{1, 2, 2, 1}, []int{2, 2})
}
