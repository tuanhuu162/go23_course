package main

import (
	"fmt"
)

func dfs(rectangles *[][]int, i int, j int) int {
	if i < 0 || j < 0 || i >= len(*rectangles) || j >= len((*rectangles)[0]) || (*rectangles)[i][j] < 1 {
		return 0
	}
	(*rectangles)[i][j] = -1
	dfs(rectangles, i+1, j)
	dfs(rectangles, i-1, j)
	dfs(rectangles, i, j-1)
	dfs(rectangles, i, j+1)
	return 1
}

func countRectangles(rectangles [][]int) int {
	m, n := len(rectangles), len(rectangles[0])

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rectangles[i][j] == 1 {
				count += dfs(&rectangles, i, j)
				for k := 0; k < m; k++ {
					fmt.Println(rectangles[k])
				}
			}
		}
	}
	return count
}

func PublicCountRectangles(rectangles [][]int) int {
	return countRectangles(rectangles)
}

func main() {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}
	count := countRectangles(arr)
	fmt.Println(count)
}
