package leetcode

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int

	// 直接使用内部函数实现回溯
	// 简化外部函数实现时需要传递多个参数的方式
	var backtrack func(cur, rest int)
	backtrack = func(cur, rest int) {
		if len(path) == k && rest == 0 {
			res = append(res, append([]int{}, path...))
			return
		}

		// 剪枝:
		// 1. 当前选取的数字个数即使再怎么加其他数字，也不可能组成目标数字 n (小于关系)
		// 2. 当前选取的数字总和已经超过目标数字 n (大于关系)
		if len(path)+10-cur < k || rest < 0 {
			return
		}

		// 跳过当前数字
		backtrack(cur+1, rest)

		// 选择当前数字
		path = append(path, cur)
		backtrack(cur+1, rest-cur)
		path = path[:len(path)-1]
	}

	backtrack(1, n)

	return res
}
