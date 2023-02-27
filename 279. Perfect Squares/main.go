package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	var c int
	for n != 0 {
		n = n - maxsquare(n)

		c++

		//fmt.Println(n)
	}
	return c
}

func maxsquare(n int) int {
	return int(math.Pow(math.Floor(math.Sqrt(float64(n))), 2))
}

func main() {
	fmt.Println(numSquares(13))
}
