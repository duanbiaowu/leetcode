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

func numIslands2(grid [][]byte) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				// 如果找到一块陆地，以该坐标为中心，上下左右四个方向继续探索
				res++
				bfs(grid, i, j)
			}
		}
	}

	return res
}

func bfs(grid [][]byte, i, j int) {
	// 将当前的起始坐标加到队列
	// 这里的队列使用二维数组表示
	queue := [][2]int{}
	queue = append(queue, [2]int{i, j})

	for len(queue) > 0 {
		// 从队列中取出第一个坐标
		i, j := queue[0][0], queue[0][1]
		// 第一个坐标出队
		queue = queue[1:]

		// 如果坐标未越界 并且 当前坐标是陆地
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == '1' {
			// 已经探索过的坐标标记为 0, 避免重复计算
			grid[i][j] = '0'

			// ⬆ 方向继续探索
			queue = append(queue, [2]int{i - 1, j})
			// ⬇ 方向继续探索
			queue = append(queue, [2]int{i + 1, j})
			// <- 方向继续探索
			queue = append(queue, [2]int{i, j - 1})
			// -> 方向继续探索
			queue = append(queue, [2]int{i, j + 1})
		}
	}
}
