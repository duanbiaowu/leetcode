package leetcode

func updateMatrix(mat [][]int) [][]int {
	var queue [][2]int

	// 将所有的 0 添加进初始队列中 (超级源点)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	dx := []int{0, 1, 0, -1} // x: 上右下左
	dy := []int{-1, 0, 1, 0} // y: 上右下左

	for len(queue) > 0 {
		for k := 0; k < 4; k++ {
			x, y := queue[0][0]+dx[k], queue[0][1]+dy[k]
			if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) && mat[x][y] == -1 {
				mat[x][y] = mat[queue[0][0]][queue[0][1]] + 1
				queue = append(queue, [2]int{x, y})
			}
		}
		queue = queue[1:]
	}

	return mat
}
