package main

import "fmt"

func isPalindrome(x int) bool {
	var reversedNum int
	tmp := x
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}
	fmt.Println(reversedNum)
	fmt.Println(x)
	return x == reversedNum
}

func main() {
	x := -121121
	fmt.Println(isPalindrome(x))
}
