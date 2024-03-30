package leetcode

func solve(board [][]byte) {
	// 题目解释中提到：任何边界上的 O 都不会被填充为 X
	// 所有的不被包围的 O 都直接或间接与边界上的 O 相连
	// 将所有可以连通的进行标记，剩下的就是无法连通的
	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])
	if cols == 0 {
		return
	}

	// 将左右边界进行标记为 "已连通"
	for row := 0; row < rows; row++ {
		dfs(board, row, 0)
		dfs(board, row, cols-1)
	}
	// 将上下边界进行标记为 "已连通"
	for col := 1; col < cols-1; col++ {
		dfs(board, 0, col)
		dfs(board, rows-1, col)
	}

	// 现在从矩阵的最外侧作为出发点
	// 上下左右四个方向已经连通
	// 只需要遍历一次矩阵
	// 将 # 修改为 0, 因为这些属于边界上的 0
	// 将 0 改为为 X, 因为这些不属于边界上的 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == '#' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, row, col int) {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] != 'O' {
		return
	}

	board[row][col] = '#'
	dfs(board, row-1, col) // ⬆  方向递归
	dfs(board, row+1, col) // ⬇  方向递归
	dfs(board, row, col-1) // <- 方向递归
	dfs(board, row, col+1) // -> 方向递归
}

func bfs(board [][]byte) {
	// 边界处理
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])
	var queue [][2]int

	// 将左右边界进行标记为 "已连通"
	for row := 0; row < rows; row++ {
		if board[row][0] == 'O' {
			// 将值为 'O' 的坐标加入队列
			queue = append(queue, [2]int{row, 0})
			board[row][0] = '#'
		}
		if board[row][cols-1] == 'O' {
			// 将值为 'O' 的坐标加入队列
			queue = append(queue, [2]int{row, cols - 1})
			board[row][cols-1] = '#'
		}
	}

	// 将上下边界进行标记为 "已连通"
	for col := 1; col < cols-1; col++ {
		if board[0][col] == 'O' {
			// 将值为 'O' 的坐标加入队列
			queue = append(queue, [2]int{0, col})
			board[0][col] = '#'
		}
		if board[rows-1][col] == 'O' {
			// 将值为 'O' 的坐标加入队列
			queue = append(queue, [2]int{rows - 1, col})
			board[rows-1][col] = '#'
		}
	}

	dx := []int{0, 0, -1, 1} // x: 上下左右
	dy := []int{-1, 1, 0, 0} // y: 上下左右

	for len(queue) > 0 {
		// 从队列中取出第一个坐标
		top := queue[0]
		// 第一个坐标出队
		queue = queue[1:]

		// 以该坐标为中心，上下左右四个方向继续探索
		for i := 0; i < 4; i++ {
			x, y := top[0]+dx[i], top[1]+dy[i]
			if x > 0 && x < rows && y > 0 && y < cols && board[x][y] == 'O' {
				queue = append(queue, [2]int{x, y})
				board[x][y] = '#'
			}
		}
	}

	// 现在从矩阵的最外侧作为出发点
	// 上下左右四个方向已经连通
	// 只需要遍历一次矩阵
	// 将 # 修改为 0, 因为这些属于边界上的 0
	// 将 0 改为为 X, 因为这些不属于边界上的 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == '#' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}
