package leetcode

func permute(nums []int) [][]int {
	var res [][]int
	var path []int

	// 初始化一个数组作为 Map
	// 用来标记递归过程中已经存在于 “当前路径” 中的元素索引
	// 避免单个元素被重复使用多次
	// 例如正常情况下，[1, 2, 3] [1, 3, 2] 都会正常组合
	// 但是 [1, 1, 1], [2, 2, 2] 属于元素重复使用，不能算作组合
	var visited = make([]bool, len(nums))

	backtrack(nums, &path, &visited, &res)

	return res
}

func backtrack(nums []int, path *[]int, visited *[]bool, res *[][]int) {
	// 当前路径的长度是否等于参数数组的长度
	// 正好满足一个组合，将当前路径加入到结果集中
	if len(*path) == len(nums) {
		// 注意这里要拷贝当前路径中的元素
		// 因为当前路径变量 path 会在递归过程中发生变化
		row := make([]int, len(nums))
		copy(row, *path)

		// 将拷贝后的路径添加到结果集中
		*res = append(*res, row)
		return
	}

	// 遍历数组，执行递归 + 回溯操作
	for i, v := range nums {
		// 剪枝策略
		// 如果当前元素未访问
		if !(*visited)[i] {
			// 将当前元素添加到 当前路径 中，形成新的当前路径
			(*visited)[i] = true
			// 将当前元素标记为已访问
			*path = append(*path, v)

			// 进入下一层递归
			backtrack(nums, path, visited, res)

			// 将当前元素标记为未访问
			(*visited)[i] = false
			// 从当前路径中删除当前元素
			// 回溯到上一步，尝试更多可能的解
			*path = (*path)[:len(*path)-1]
		}
	}
}

func permuteDFS(nums []int) [][]int {
	var res [][]int

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums)-1 {
			res = append(res, append([]int{}, nums...))
			return
		}

		for i := index; i < len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			dfs(index + 1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}

	dfs(0)

	return res
}
