package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	res, left := 0, 0
	for i := range s {
		if _, ok := mp[s[i]]; ok && mp[s[i]] >= left {
			left = mp[s[i]] + 1
		} else if i-left+1 > res {
			res = i - left + 1
		}
		mp[s[i]] = i
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwekw"))
}
