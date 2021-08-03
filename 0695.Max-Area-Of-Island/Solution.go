package leetcode

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(grid, i, j))
			}
		}
	}
	return res
}

func dfs(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}

	grid[x][y] = 0
	res := 1
	res += dfs(grid, x+1, y)
	res += dfs(grid, x-1, y)
	res += dfs(grid, x, y-1)
	res += dfs(grid, x, y+1)

	return res
}

func bfs(grid [][]int) int {
	res := 0
	var queue [][]int

	dx := []int{0, 1, 0, -1} // x: 上右下左
	dy := []int{-1, 0, 1, 0} // y: 上右下左

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cur := 0
			queue = append(queue, []int{i, j})

			for len(queue) > 0 {
				x, y := queue[0][0], queue[0][1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					cur++
					grid[x][y] = 0

					for k := 0; k < 4; k++ {
						queue = append(queue, []int{x + dx[k], y + dy[k]})
					}
				}
				queue = queue[1:]
			}
			res = max(res, cur)
		}
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
