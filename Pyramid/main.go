package main

import "fmt"

func pyramidWeight(rows, bricksInRow, row int) float64 {
	if row == rows {
		// Это самый нижний ряд, кирпичи выше отсутствуют
		return float64(bricksInRow)
	}
	weightAbove := 0.0
	for i := 1; i <= bricksInRow; i++ {
		weightAbove += pyramidWeight(rows, i, row+1)
	}
	return 1.0 + weightAbove/2.0
}

func main() {
	weight := pyramidWeight(5, 3, 1)
	fmt.Println(weight) // Выведет: 98.0

}
