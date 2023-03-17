package main

import "strings"

type Trie struct {
	List []string
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	this.List = append(this.List, word)
}

func (this *Trie) Search(word string) bool {
	for _, item := range this.List {
		if item == word {
			return true
		}
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, item := range this.List {
		if strings.Index(item, prefix) == 0 {
			return true
		}
	}

	return false
}

func main() {

}
