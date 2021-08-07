package leetcode

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	if m == 0 {
		return mat
	}
	n := len(mat[0])

	var res = make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	var seen = make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}

	var queue [][2]int

	// 将所有的 0 添加进初始队列中
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				seen[i][j] = true
			}
		}
	}

	dx := []int{0, 1, 0, -1} // x: 上右下左
	dy := []int{-1, 0, 1, 0} // y: 上右下左

	// BFS
	for len(queue) > 0 {
		for k := 0; k < 4; k++ {
			x, y := queue[0][0]+dx[k], queue[0][1]+dy[k]
			if x >= 0 && x < m && y >= 0 && y < n && !seen[x][y] {
				res[x][y] = mat[queue[0][0]][queue[0][1]] + 1
				queue = append(queue, [2]int{x, y})
				seen[x][y] = true
			}
		}

		queue = queue[1:]
	}

	return res
}
