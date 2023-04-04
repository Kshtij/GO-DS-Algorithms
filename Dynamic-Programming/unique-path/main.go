package main

import "log"

func main() {
	log.Println("Unique Paths :", uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)

	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		grid[i][n-1] = 1
	}

	for j := 0; j < n; j++ {
		grid[m-1][j] = 1
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			grid[i][j] = grid[i][j+1] + grid[i+1][j]
		}
	}

	return grid[0][0]
}
