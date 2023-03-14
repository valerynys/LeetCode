package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Fields(s)
	sort.Slice(words, func(i, j int) bool {
		return words[i][len(words[i])-1] < words[j][len(words[j])-1]
	})
	lastSubstrings := make([]string, len(words))

	for i, word := range words {
		lastSubstrings[i] = word[:len(word)-1]
	}
	str := strings.Join(lastSubstrings, " ")
	return str
}

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}
