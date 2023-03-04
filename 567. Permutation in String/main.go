package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	need := [26]int{}
	for i := range s1 {
		need[s1[i]-'a']++
	}
	fmt.Println(need)
	have := [26]int{}
	for r := range s2 {
		if r >= len(s1) {
			//fmt.Println(r)
			fmt.Println(s2[r] - 'a')
			have[s2[r-len(s1)]-'a']--
		}
		have[s2[r]-'a']++
		if need == have {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("abe", "eidbaoo"))
}
