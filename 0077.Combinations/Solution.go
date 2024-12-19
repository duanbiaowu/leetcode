package leetcode

func combine(n int, k int) [][]int {
	if k <= 0 || n < k {
		return [][]int{}
	}

	var path []int
	var res [][]int
	backtrack(n, k, 1, &path, &res)
	return res
}

func backtrack(n, k, begin int, path *[]int, res *[][]int) {
	// 剪枝：path 长度加上区间 [begin, n] 的长度小于 k， 不可能构造出长度为 k 的 组合
	if len(*path)+(n-begin+1) < k {
		return
	}
	if len(*path) == k {
		comb := make([]int, k)
		copy(comb, *path)
		*res = append(*res, comb)
		// *res = append(*res, append([]int{}, (*path)...))
		return
	}
	for i := begin; i <= n; i++ {
		*path = append(*path, i)
		backtrack(n, k, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
