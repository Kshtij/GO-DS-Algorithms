package main

func main() {

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}

	if obstacleGrid[m-1][n-1] != -1 {
		obstacleGrid[m-1][n-1] = 1
	}

	for i := m - 2; i >= 0; i-- {
		if obstacleGrid[i][n-1] != -1 && obstacleGrid[i+1][n-1] != -1 {
			obstacleGrid[i][n-1] = 1
		} else {
			obstacleGrid[i][n-1] = -1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if obstacleGrid[m-1][i] != -1 && obstacleGrid[m-1][i+1] != -1 {
			obstacleGrid[m-1][i] = 1
		} else {
			obstacleGrid[m-1][i] = -1
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] != -1 {
				var num1, num2 int
				if obstacleGrid[i][j+1] != -1 {
					num1 = obstacleGrid[i][j+1]
				}

				if obstacleGrid[i+1][j] != -1 {
					num2 = obstacleGrid[i+1][j]
				}

				obstacleGrid[i][j] = num1 + num2

			}
		}
	}

	if obstacleGrid[0][0] == -1 {
		return 0
	}
	return obstacleGrid[0][0]
}
