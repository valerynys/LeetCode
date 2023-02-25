package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	i, j, d := 0, 0, 0
	for k := 1; k <= n*n; k++ {
		res[i][j] = k

		di, dj := dirs[d%4][0], dirs[d%4][1]
		if i+di < 0 || i+di >= n || j+dj < 0 || j+dj >= n || res[i+di][j+dj] != 0 {
			d++
			di, dj = dirs[d%4][0], dirs[d%4][1]
		}

		i, j = i+di, j+dj
	}

	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}
