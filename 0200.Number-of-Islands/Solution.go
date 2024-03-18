package leetcode

func numIslands(grid [][]byte) int {
	res := 0

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				// 如果找到一块陆地，以该坐标为中心，上下左右四个方向继续探索
				res++
				dfs(grid, row, col)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, row, col int) {
	// 如果坐标已越界或者当前坐标不是陆地
	// 直接返回
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}

	// 已经探索过的坐标标记为 0, 避免重复计算
	grid[row][col] = '0'
	dfs(grid, row-1, col) // ⬆  方向递归
	dfs(grid, row+1, col) // ⬇  方向递归
	dfs(grid, row, col-1) // <- 方向递归
	dfs(grid, row, col+1) // -> 方向递归
}
