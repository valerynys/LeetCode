package main

import (
	"fmt"
)

func removeEmoticons(s string) string {
	var res []rune
	open, close := 0, 0
	for _, c := range s {
		switch c {
		case ':':
			if open == 0 && close == 0 {
				res = append(res, c)
			} else {
				res = append(res, ':')
			}
		case '-':
			if len(res) > 0 && res[len(res)-1] == ':' {
				res = append(res, '-')
			} else {
				res = append(res, c)
			}
		case ')':
			if len(res) > 0 && res[len(res)-1] == '-' {
				close++
				if close > open {
					res = res[:len(res)-open-close-1]
					open, close = 0, 0
				}
			} else {
				res = append(res, c)
				open, close = 0, 0
			}
		case '(':
			if len(res) > 0 && res[len(res)-1] == '-' {
				open++
				if open > close {
					res = res[:len(res)-open-close-1]
					open, close = 0, 0
				}
			} else {
				res = append(res, c)
				open, close = 0, 0
			}
		default:
			res = append(res, c)
			open, close = 0, 0
		}
	}
	return string(res)
}

func main() {
	s1 := "I work in google :-)))"
	s2 := "Cool :-) and I failed assesment there:-(("
	s3 := "lol:)"
	s4 := "YEEEEE BOIIII!!!! :-))(())"
	s5 := "Cringe :-)))))))))))))))"

	fmt.Println(removeEmoticons(s1)) // "I work in google "
	fmt.Println(removeEmoticons(s2)) // "Cool  and I failed assesment there"
	fmt.Println(removeEmoticons(s3)) // "lol:)"
	fmt.Println(removeEmoticons(s4)) // "YEEEEE BOIIII!!!! (())"
	fmt.Println(removeEmoticons(s5)) // "Cringe "

}
