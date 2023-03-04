package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i < n+1; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			res = append(res, "FizzBuzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		case i%5 == 0:
			res = append(res, "Buzz")
		default:
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
func main() {
	fmt.Println(fizzBuzz(20))
}
