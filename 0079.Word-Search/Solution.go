package leetcode

func exist(board [][]byte, word string) bool {
	n := len(board)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(board, word, visited, 0, i, j) {
				return true
			}
		}
	}

	return false
}

var directions = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func inBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func backtrack(board [][]byte, word string, visited [][]bool, begin, x, y int) bool {
	if begin == len(word)-1 {
		return board[x][y] == word[begin]
	}

	if board[x][y] == word[begin] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + directions[i][0]
			ny := y + directions[i][1]
			if inBoard(board, nx, ny) && !visited[nx][ny] && backtrack(board, word, visited, begin+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}

	return false
}
