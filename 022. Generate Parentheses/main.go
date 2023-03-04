package main

import "fmt"

func generateParenthesis(n int) []string {
	return generateParenthesisHelper("", 0, n)
}

func generateParenthesisHelper(s string, open, n int) []string {
	if n == 0 && open == 0 {
		return []string{s}
	}
	result := []string{}
	if open > 0 {
		result = append(result, generateParenthesisHelper(s+")", open-1, n)...)
	}
	if n > 0 {
		result = append(result, generateParenthesisHelper(s+"(", open+1, n-1)...)
	}

	return result
}

func main() {
	fmt.Println(generateParenthesis(5))
}
