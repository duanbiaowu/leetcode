package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var path []int
	var res [][]int
	sort.Ints(candidates)

	backtrack(candidates, target, 0, &path, &res)

	return res
}

func backtrack(candidates []int, target, begin int, path *[]int, res *[][]int) {
	// 回溯公式: 查找结果是什么，终止条件就是什么
	if target == 0 {
		*res = append(*res, append([]int{}, *path...))
		return
	}

	for i := begin; i < len(candidates); i++ {
		// 剪枝
		if candidates[i] > target {
			return
		}
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}

		// candidates 中的 同一个 数字不可以重复
		// i + 1 作为起点，查找 target-candidates[i]
		// 和 39 题最主要的差异点
		*path = append(*path, candidates[i])
		backtrack(candidates, target-candidates[i], i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
