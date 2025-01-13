package leetcode

import leetcode "leetcode/0040.Combination-Sum-II"

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int

	backtrack(k, 1, n, &path, &res)

	return res
}

func backtrack(k, start, target int, path *[]int, res *[][]int) {
	// 回溯结束条件
	if len(*path) == k && target == 0 {
		*res = append(*res, append([]int{}, *path...))
		return
	}

	// 剪枝条件
	// 1. 当前选取的数字总和已经大于 n (所以 target 为负数)
	// 2. 当前选择的数字总和太小了, 即使再添加其他数字，长度也无法满足要求 (小于 k)
	//     所以没有继续回溯的必要了
	if target < 0 || len(*path)-start+10 < k {
		return
	}

	// 不选择当前数字
	backtrack(k, start+1, target, path, res)

	// 选择当前数字
	*path = append(*path, start)
	backtrack(k, start+1, target-start, path, res)
	*path = (*path)[:len(*path)-1]
}

// 也可以直接复用 combinationSum2 方法
// 首先计算出数字 1 - 9 的所有 sum 等于 n 组合
// 然后过滤掉长度不等于 k 的子集
func combinationSum3Simple(k int, n int) [][]int {
	nums := make([]int, 9)
	for i := range nums {
		nums[i] = i + 1
	}

	list := leetcode.CombinationSum2(nums, n)

	var res [][]int

	for _, row := range list {
		if len(row) == k {
			res = append(res, row)
		}
	}

	return res
}
