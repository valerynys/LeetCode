package main

import "fmt"

func tribonacci(n int) int {
	n1 := 0
	n2 := 1
	n3 := 1
	temp1 := 0
	temp2 := 0
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}

	for i := 3; i < n; i++ {
		temp1 = n1
		temp2 = n2
		n1 = n2
		n2 = n3
		n3 = temp1 + temp2 + n3
	}
	return n1 + n2 + n3
}

func main() {
	fmt.Println(tribonacci(50))
}
