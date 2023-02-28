package main

import "fmt"

func categorizeBox(length int, width int, height int, mass int) string {
	bulky := length >= 10000 || width >= 10000 || height >= 10000 || length*width*height >= 1000000000
	heavy := mass >= 100
	if bulky != heavy {
		if bulky {
			return "Bulky"
		}
		return "Heavy"
	}
	if bulky {
		return "Both"
	}
	return "Neither"
}

func main() {
	fmt.Println(categorizeBox(1000, 35, 700, 300))
}
