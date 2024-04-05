package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	var path []int
	var res [][]int

	// 为了在递归过程中优化剪枝
	// 提前对数组进行排序
	sort.Ints(candidates)

	backtrack(candidates, target, 0, &path, &res)

	return res
}

func backtrack(candidates []int, target, begin int, path *[]int, res *[][]int) {
	// 如果当前查找目标值等于 0
	// 正好满足一个组合，将当前路径加入到结果集中
	if target == 0 {
		// 这里使用了另外一种语法来实现追加功能
		*res = append(*res, append([]int{}, *path...))
		return
	}

	for i := begin; i < len(candidates); i++ {
		// 剪枝 (前提：已排序)
		// 数组排序之后
		// 如果当前元素大于目标值，那么当前元素之后的所有元素都不可能出现有效的组合了
		// 例如数组排序之后为 [2,3,6,7], target = 5
		// 那么当前元素遍历到 6 时，就可以直接剪枝了
		if candidates[i] > target {
			return
		}

		// 将当前元素添加到 当前路径 中，形成新的当前路径
		*path = append(*path, candidates[i])

		// candidates 中的 同一个 数字可以 无限制重复被选取
		// i 作为起点，查找 target-candidates[i]
		// 指定起始位置为当前位置
		backtrack(candidates, target-candidates[i], i, path, res)

		// 从当前路径中删除当前元素
		// 回溯到上一步，尝试更多可能的解
		*path = (*path)[:len(*path)-1]
	}
}

// 实现略微臃肿
// 递归过程中传递 “当前累积和 sum”
func backtrack2(candidates []int, target, sum, begin int, path *[]int, res *[][]int) {
	// 剪枝
	if sum > target {
		return
	}
	if sum == target {
		*res = append(*res, append([]int{}, *path...))
		return
	}

	for i := begin; i < len(candidates); i++ {
		// 剪枝
		if candidates[i] > target {
			return
		}

		*path = append(*path, candidates[i])
		backtrack2(candidates, target, sum+candidates[i], i, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
