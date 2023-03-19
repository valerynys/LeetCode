package main

import "fmt"

type TrieNode struct {
	isEndOfWord bool
	children    map[byte]*TrieNode
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{isEndOfWord: false, children: make(map[byte]*TrieNode)}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{isEndOfWord: false, children: make(map[byte]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchHelper(word, 0, this.root)
}

func (this *WordDictionary) searchHelper(word string, index int, node *TrieNode) bool {
	if index == len(word) {
		return node.isEndOfWord
	}
	char := word[index]
	if char == '.' {
		for _, child := range node.children {
			if this.searchHelper(word, index+1, child) {
				return true
			}
		}
		return false
	}
	child, ok := node.children[char]
	if !ok {
		return false
	}
	return this.searchHelper(word, index+1, child)
}

func main() {
	dict := Constructor()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")
	fmt.Println(dict.Search("pad")) // false
	fmt.Println(dict.Search("bad")) // true
	fmt.Println(dict.Search(".ad")) // true
	fmt.Println(dict.Search("b..")) // true
}
