package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	start int
	end   int
}

type IntervalList []Interval

func (il IntervalList) Len() int {
	return len(il)
}

func (il IntervalList) Less(i, j int) bool {
	return il[i].start < il[j].start
}

func (il IntervalList) Swap(i, j int) {
	il[i], il[j] = il[j], il[i]
}

func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(IntervalList(intervals))

	merged := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := &merged[len(merged)-1]
		if current.start <= lastMerged.end {
			if current.end > lastMerged.end {
				lastMerged.end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func main() {
	intervals := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := MergeIntervals(intervals)
	fmt.Println(merged)
}
