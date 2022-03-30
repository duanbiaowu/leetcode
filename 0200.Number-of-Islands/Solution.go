package leetcode

func numIslands(grid [][]byte) int {
	cnt := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				cnt++
				dfs(grid, i, j)
			}
		}
	}

	return cnt
}

func dfs(grid [][]byte, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
}
