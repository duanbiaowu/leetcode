package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var path []int
	var res [][]int
	backtrack(candidates, target, 0, &path, &res)
	return res
}

func backtrack(candidates []int, target, begin int, path *[]int, res *[][]int) {
	// 回溯公式: 查找结果是什么，终止条件就是什么
	if target == 0 {
		row := make([]int, len(*path))
		copy(row, *path)
		*res = append(*res, row)
		return
	}

	for i := begin; i < len(candidates); i++ {
		if candidates[i] <= target {
			if i > begin && candidates[i] == candidates[i-1] {
				continue
			}
			*path = append(*path, candidates[i])
			backtrack(candidates, target-candidates[i], i+1, path, res)
			*path = (*path)[:len(*path)-1]
		}
	}
}
