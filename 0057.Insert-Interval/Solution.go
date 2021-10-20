package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	if newInterval == nil {
		return intervals
	}
	if intervals == nil {
		return append(make([][]int, 0), newInterval)
	}

	var res [][]int
	left, right := newInterval[0], newInterval[1]
	merged := false
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > right {
			// 在当前区间右侧并且无交集
			if !merged {
				merged = true
				res = append(res, []int{left, right})
			}
			res = append(res, intervals[i])
		} else if intervals[i][1] < left {
			// 在当前区间左侧并且无交集
			res = append(res, intervals[i])
		} else {
			// 与当前区间有交集，计算并集
			left = min(left, intervals[i][0])
			right = max(right, intervals[i][1])
		}
	}
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
