package main

import "log"

func main() {
	log.Println()
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := m - 2; i >= 0; i-- {
		grid[i][n-1] = grid[i][n-1] + grid[i+1][n-1]
	}

	for i := n - 2; i >= 0; i-- {
		grid[m-1][i] = grid[m-1][i] + grid[m-1][i+1]
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {

			if grid[i][j+1] < grid[i+1][j] {
				grid[i][j] = grid[i][j] + grid[i][j+1]
			} else {
				grid[i][j] = grid[i][j] + grid[i+1][j]
			}
		}
	}

	return grid[0][0]
}
