package main

import "fmt"

func maximumPopulation(logs [][]int) int {
	delta := make([]int, 101)
	offset := 1950
	for _, log := range logs {
		delta[log[0]-offset] += 1
		delta[log[1]-offset] -= 1
	}
	mx, res, curr := 0, 0, 0
	for i, val := range delta {
		curr += val
		if curr > mx {
			mx = curr
			res = i
		}
	}
	return res + offset
}

func main() {
	fmt.Println(maximumPopulation([][]int{{1993, 1999}, {2000, 2010}}))
}
