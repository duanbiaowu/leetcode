package leetcode

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	n := len(obstacleGrid)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, len(obstacleGrid[0]))
	}

	var res [][]int
	backtrack(obstacleGrid, 0, 0, visited, &res)
	return res
}

func inBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func backtrack(board [][]int, x, y int, visited [][]bool, paths *[][]int) bool {
	if !inBoard(board, x, y) || board[x][y] == 1 || visited[x][y] {
		return false
	}

	*paths = append(*paths, []int{x, y})
	if x == len(board)-1 && y == len(board[0])-1 {
		return true
	}

	visited[x][y] = true
	if backtrack(board, x+1, y, visited, paths) || backtrack(board, x, y+1, visited, paths) {
		return true
	}

	*paths = (*paths)[:len(*paths)-1]
	return false
}
