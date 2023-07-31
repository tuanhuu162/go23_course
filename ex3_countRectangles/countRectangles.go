package main

import (
	"fmt"
)

func isRectangle(rectangles *[][]int, startRow int, endRow int, startCol int, endCol int) bool {
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			if (*rectangles)[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func markingRectangle(rectangles *[][]int, startRow int, endRow int, startCol int, endCol int) {
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			(*rectangles)[i][j] = 0
		}
	}
}

func countRectangles(rectangles [][]int) int {
	count := 0

	if len(rectangles) == 0 {
		return count
	}

	m, n := len(rectangles), len(rectangles[0])
	for startRow := 0; startRow < m; startRow++ {
		for startCol := 0; startCol < n; startCol++ {
			if rectangles[startRow][startCol] == 1 {
				// searching down from bigest rectangle to smallest rectangle
				for endRow := len(rectangles) - 1; endRow >= startRow; endRow-- {
					offset := 0
					for endCol := len((rectangles)[0]) - 1; endCol >= startCol; endCol-- {
						if isRectangle(&rectangles, startRow, endRow, startCol, endCol) {
							count++
							markingRectangle(&rectangles, startRow, endRow, startCol, endCol)
							// we will stop searching by assign the indice to the stop condition
							offset, endRow, endCol = endCol-startCol+1, startRow, startCol
							// fmt.Println(startRow, startCol, endRow, endCol, count)
							// for k := 0; k < m; k++ {
							// 	fmt.Println(rectangles[k])
							// }
						}
						// then ignore the marked rectangle and continue searching
						startCol += offset
					}
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
