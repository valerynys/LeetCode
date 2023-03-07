package main

import "fmt"

func minimumTime(time []int, totalTrips int) int64 {
	var l, r int
	for i := range time {
		if r < time[i] {
			r = time[i]
		}
	}
	r *= totalTrips

	for l <= r {
		mid := (l + r) / 2
		if check(time, totalTrips, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return int64(l)
}

func check(time []int, totalTrips int, try int) bool {
	var total int

	for i := range time {
		total += try / time[i]
	}

	return total >= totalTrips
}

func main() {
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
}
