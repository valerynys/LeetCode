package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		p, q := firstList[i], secondList[j]

		a, b := max(p[0], q[0]), min(p[1], q[1])
		if b >= a {
			ans = append(ans, []int{a, b})
		}

		if p[1] > q[1] {
			j += 1
		} else {
			i += 1
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	first := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	second := [][]int{{2, 4}, {1, 5}, {9, 14}, {21, 27}}
	fmt.Println(intervalIntersection(first, second))
}
