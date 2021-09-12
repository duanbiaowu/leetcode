package leetcode

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	// 题目解释中提到：任何边界上的 O 都不会被填充为 X
	// 所有的不被包围的 O 都直接或间接与边界上的 O 相连
	n, m := len(board), len(board[0])

	// 左右边界
	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}
	// 上下边界
	for i := 1; i < m-1; i++ {
		dfs(board, 0, i)
		dfs(board, n-1, i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	n, m := len(board), len(board[0])
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}

	board[x][y] = '#'
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}
