package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	count := 1
	i := 0
	for j := 1; j <= len(chars); j++ {
		if j < len(chars) && chars[j] == chars[j-1] {
			count++
		} else {
			chars[i] = chars[j-1]
			i++
			if count > 1 {
				for _, c := range strconv.Itoa(count) {
					chars[i] = byte(c)
					i++
				}
			}
			count = 1
		}
	}
	return i
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}
