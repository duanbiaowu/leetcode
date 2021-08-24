package leetcode

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	nr, nc := len(grid), len(grid[0])
	count := 0
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, r, c int) {
	nr, nc := len(grid), len(grid[0])
	if r < 0 || r >= nr || c < 0 || c >= nc || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
}
