package leetcode

func movingCount(m int, n int, k int) int {
	if m <= 0 || n <= 0 || k < 0 {
		return 0
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(m, n, k, 0, 0, visited)
}

func dfs(m, n, k, i, j int, visited [][]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || (i/10+i%10+j/10+j%10) > k {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(m, n, k, i+1, j, visited) +
		dfs(m, n, k, i-1, j, visited) +
		dfs(m, n, k, i, j+1, visited) +
		dfs(m, n, k, i, j-1, visited)
}
