package main

import "log"

func main() {
	grid := [][]byte{[]byte{'1', '1', '1', '1', '0'}, []byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '0', '0', '0'}}
	log.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	return islandCalc(grid)
}

type rowCol struct {
	row int
	col int
}

func islandCalc(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	isLandCount := 0

	visited := make(map[rowCol]interface{})

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rc := rowCol{
				row: i,
				col: j,
			}
			_, ok := visited[rc]
			if grid[rc.row][rc.col] == '1' && !ok {
				isLandCount++
				bfs(rc, rows, cols, visited, grid)
			}

		}
	}
	return isLandCount

}

func bfs(rc rowCol, numRows, numCols int, visited map[rowCol]interface{}, grid [][]byte) {
	//log.Println("call 1")
	queue := make([]rowCol, 0, 1)
	queue = append(queue, rc)

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		neighbours := getNeighbours(node, numRows, numCols)
		for _, neighbour := range neighbours {
			_, ok := visited[neighbour]
			if grid[neighbour.row][neighbour.col] == '1' && !ok {
				//log.Println("appending neighbour :", neighbour)
				queue = append(queue, neighbour)
			}

		}
		visited[node] = nil
		//log.Println("visited :", node)
	}
	//log.Println("visited set :", visited)

}

func getNeighbours(rc rowCol, rows, cols int) []rowCol {
	retList := make([]rowCol, 0, 1)
	if rc.col-1 >= 0 {
		retList = append(retList, rowCol{rc.row, rc.col - 1})
	}
	if rc.row-1 >= 0 {
		retList = append(retList, rowCol{rc.row - 1, rc.col})
	}

	if rc.col+1 < cols {
		retList = append(retList, rowCol{rc.row, rc.col + 1})
	}
	if rc.row+1 < rows {
		retList = append(retList, rowCol{rc.row + 1, rc.col})
	}

	return retList
}
