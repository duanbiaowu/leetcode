package leetcode

var directions = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// 检测指定坐标是否位于表格内
func inBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func exist(board [][]byte, word string) bool {
	// 表示指定坐标在本轮回溯过程中是否已经访问过
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	// 遍历矩阵的每个坐标
	// start 参数都是从 0 开始
	// 表示以当前坐标为起点，从参数 word 字符串的第一个字符开始匹配
	// 本质上依然是一种暴力搜索方法
	for i := range board {
		for j := range board[0] {
			if backtrack(board, word, visited, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func backtrack(board [][]byte, word string, visited [][]bool, start, x, y int) bool {
	// 剪枝条件: 当前匹配长度等于目标参数 word 字符串长度
	// 直接对比当前坐标的字符是否等于 word 字符串的最后一个字符即可
	if start == len(word)-1 {
		return board[x][y] == word[start]
	}

	if board[x][y] == word[start] {
		visited[x][y] = true

		// 分别以当前坐标的上下左右 4 个坐标为新的坐标
		// 开始新一轮的回溯过程
		for i := 0; i < 4; i++ {
			newX := x + directions[i][0]
			newY := y + directions[i][1]
			if inBoard(board, newX, newY) && !visited[newX][newY] &&
				backtrack(board, word, visited, start+1, newX, newY) {
				return true
			}
		}

		visited[x][y] = false
	}

	return false
}
