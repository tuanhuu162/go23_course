package rectangles

func dfs(rectangles *[][]int, i int, j int) int {
	if i < 0 || j < 0 || i >= len(*rectangles) || j >= len((*rectangles)[0]) || (*rectangles)[i][j] < 1 {
		return 0
	}
	(*rectangles)[i][j] = -1
	dfs(rectangles, i+1, j)
	dfs(rectangles, i, j+1)
	return 1
}

func countRectangles(rectangles [][]int) int {
	count := 0

	if len(rectangles) == 0 {
		return count
	}

	m, n := len(rectangles), len(rectangles[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rectangles[i][j] == 1 {
				count += dfs(&rectangles, i, j)
				// for k := 0; k < m; k++ {
				// 	fmt.Println(rectangles[k])
				// }
			}
		}
	}
	return count
}

func PublicCountRectangles(rectangles [][]int) int {
	return countRectangles(rectangles)
}
