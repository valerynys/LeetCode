package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	i := 0
	for i < len(flowerbed) {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
		if count >= n {
			return true
		}
		i++
	}
	return false
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))

}