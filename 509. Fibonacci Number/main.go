package main

import "fmt"

func fib(n int) int {
	n1 := 0
	n2 := 1
	temp := 0
	if n == 0 {
		return 0
	}
	for i := 2; i < n; i++ {
		temp = n1
		n1 = n2
		n2 = temp + n2
	}
	return n1 + n2
}

func main() {
	fmt.Println(fib(10))
}
