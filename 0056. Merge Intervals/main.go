package main

import (
	"fmt"
	"sort"
)

// Time: O(nlogn + n) = O(nlogn)
// Space: O(n)
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Time: O(nlogn)
	sortIntervals(intervals)

	// Space: O(n)
	mergedIntervals := make([][]int, 0, len(intervals))
	mergedIntervals = append(mergedIntervals, intervals[0])

	// Time: O(n)
	for _, interval := range intervals[1:] {
		if top := mergedIntervals[len(mergedIntervals)-1]; interval[0] > top[1] {
			mergedIntervals = append(mergedIntervals, interval)
		} else if interval[1] > top[1] {
			top[1] = interval[1]
		}
	}

	return mergedIntervals
}

// If an interval is {x, y},
// sortIntervals sorts the intervals by their "x" value.
// Time: O(nlogn)
// Space: O(1)
func sortIntervals(intervals [][]int) {
	// Time: O(nlogn)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
}

func main() {
	intervals := [][]int{{1, 3}, {2, 4}, {1, 2}, {15, 18}}
	merged := merge(intervals)
	fmt.Println(merged)
}
