package main

import "fmt"

func dfs(grid [][]byte, row int, col int) {
	rows := len(grid)
	cols := len(grid[0])
	grid[row][col] = '2'
	if row > 0 && grid[row-1][col] == '1' {
		dfs(grid, row-1, col)
	}
	if col > 0 && grid[row][col-1] == '1' {
		dfs(grid, row, col-1)
	}
	if row < rows-1 && grid[row+1][col] == '1' {
		dfs(grid, row+1, col)
	}
	if col < cols-1 && grid[row][col+1] == '1' {
		dfs(grid, row, col+1)
	}
}

func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	c := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				c++
				dfs(grid, i, j)
			}
		}
	}
	return c
}
func main() {
	arr := [][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(arr))
}
