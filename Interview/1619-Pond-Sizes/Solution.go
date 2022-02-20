package leetcode

import "sort"

func pondSizes(land [][]int) []int {
	var res []int
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 0 {
				res = append(res, dfs(land, i, j))
			}
		}
	}
	sort.Ints(res)
	return res
}

func dfs(land [][]int, i, j int) int {
	if i < 0 || i >= len(land) || j < 0 || j >= len(land[0]) || land[i][j] != 0 {
		return 0
	}
	land[i][j] = -1
	cnt := 1
	cnt += dfs(land, i+1, j)
	cnt += dfs(land, i-1, j)
	cnt += dfs(land, i, j+1)
	cnt += dfs(land, i, j-1)
	cnt += dfs(land, i+1, j+1)
	cnt += dfs(land, i-1, j+1)
	cnt += dfs(land, i-1, j-1)
	cnt += dfs(land, i+1, j-1)
	return cnt
}
