package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	// 边界情况处理
	if newInterval == nil {
		return intervals
	}
	if intervals == nil {
		return [][]int{newInterval}
	}

	var res [][]int
	left, right := newInterval[0], newInterval[1]
	// 区间是否已合并标记
	merged := false

	for _, val := range intervals {
		if val[0] > right {
			// 新区间和当前区间右边界无重叠
			if !merged {
				// 更新区间是否已合并标记
				merged = true
				res = append(res, []int{left, right})
			}
			res = append(res, val)
		} else if val[1] < left {
			// 新区间和当前区间左边界无重叠
			res = append(res, val)
		} else {
			// 新区间和当前区间重叠
			left = min(left, val[0])
			right = max(right, val[1])
		}
	}

	// 代码执行到这里还没有合并
	// 说明新区间比当前所有区间都要大
	if !merged {
		res = append(res, []int{left, right})
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
