package main

import "fmt"

type BrowserHistory struct {
	urls      []string
	currIndex int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (this *BrowserHistory) Visit(url string) {
	this.currIndex++
	if this.currIndex < len(this.urls) {
		this.urls[this.currIndex] = url
	} else {
		this.urls = append(this.urls, url)
	}
	this.urls = this.urls[:this.currIndex+1]
}

func (this *BrowserHistory) Back(steps int) string {
	if steps > this.currIndex {
		this.currIndex = 0
	} else {
		this.currIndex -= steps
	}
	return this.urls[this.currIndex]
}

func (this *BrowserHistory) Forward(steps int) string {
	if steps > len(this.urls)-1-this.currIndex {
		this.currIndex = len(this.urls) - 1
	} else {
		this.currIndex += steps
	}
	return this.urls[this.currIndex]
}

func main() {
	history := Constructor("https://www.google.com")

	// Visit some websites
	history.Visit("https://www.wikipedia.org")
	history.Visit("https://www.nytimes.com")
	history.Visit("https://www.github.com")

	// Move back and forward in history
	fmt.Println(history.Back(2))    // should print https://www.wikipedia.org
	fmt.Println(history.Forward(1)) // should print https://www.nytimes.com

	// Visit a new website and move back again
	history.Visit("https://www.stackoverflow.com")
	fmt.Println(history.Back(3)) // should print https://www.google.com
}
