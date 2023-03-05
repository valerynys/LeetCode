package main

import "fmt"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	champagne := make([][]float64, 100)
	for i := range champagne {
		champagne[i] = make([]float64, i+1)
	}
	champagne[0][0] = float64(poured)
	for i := 0; i < 99; i++ {
		for j := 0; j <= i; j++ {
			if champagne[i][j] > 1.0 {
				excess := champagne[i][j] - 1.0
				champagne[i+1][j] += excess / 2.0
				champagne[i+1][j+1] += excess / 2.0
				champagne[i][j] = 1.0
			}
		}
	}
	return champagne[query_row][query_glass]
}

func main() {
	fmt.Println(champagneTower(1, 1, 1))
}
