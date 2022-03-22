package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	var path []int
	var res [][]int
	sort.Ints(candidates)
	backtrack(candidates, target, 0, &path, &res)
	return res
}

func backtrack(candidates []int, target, begin int, path *[]int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, *path...))
		return
	}

	for i := begin; i < len(candidates); i++ {
		// 剪枝 (前提：已排序)
		if candidates[i] > target {
			return
		}

		*path = append(*path, candidates[i])
		// candidates 中的 同一个 数字可以 无限制重复被选取
		// i 作为起点，查找 target-candidates[i]
		backtrack(candidates, target-candidates[i], i, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
