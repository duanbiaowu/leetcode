package leetcode

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return intervals
	}

	// 使用标准库排序
	// sort.Slice(intervals, func(i, j int) bool {
	// 	// 只需要比较左边界的值即可
	// 	return intervals[i][0] < intervals[j][0]
	// })

	// 使用自定义排序
	quickSort(intervals, 0, n-1)

	// 第一个参数区间可以直接放入结果数组
	// 后续有重叠的情况下，直接合并即可
	res := [][]int{intervals[0]}
	pos := 0

	for i := 1; i < n; i++ {
		if intervals[i][0] > res[pos][1] {
			// 区间不重叠，增加新的区间
			// 例如: [8,10],[15,18]]
			res = append(res, intervals[i])
			pos++
		} else {
			// 合并重叠的区间
			// 虽然已经排序，但是排序的基准是区间的左边界元素
			// 合并时，区间右边界元素必须以两个区间的右侧最大值为准
			// 例如这个边界情况：[[1,4],[2,3]]
			res[pos][1] = max(res[pos][1], intervals[i][1])
		}
	}

	return res
}

func quickSort(intervals [][]int, left, right int) {
	if left < right {
		p := partition(intervals, left, right)
		quickSort(intervals, left, p-1)
		quickSort(intervals, p+1, right)
	}
}

func partition(intervals [][]int, left, right int) int {
	pivot := intervals[right]
	i, j := left, right

	for i < j {
		// 注意: 右侧的判断条件是必须的
		// intervals[i][0] == pivot[0] && intervals[i][1] < pivot[1]
		// 如果没有这个判断条件，遇到相同的区间时会陷入死循环
		// 例如 [[1,4],[1,4]]

		// for i < j && (intervals[i][0] < pivot[0] || intervals[i][0] == pivot[0] && intervals[i][1] < pivot[1]) {
		// 	i++
		// }

		// 和上面的条件同理: 右侧的判断条件也是必须的
		// 非常细微的边界：intervals[j][1] >= pivot[1]

		// for i < j && (intervals[j][0] > pivot[0] || intervals[j][0] == pivot[0] && intervals[j][1] >= pivot[1]) {
		// 	j--
		// }

		// 上面的两个条件是根据区间的左右边界进行的 “全排序”
		// 但是实际上符合题目的要求，只需要根据左边界进行排序即可
		// 因为由边界的值会在区间合并过程中不断被更新
		for i < j && intervals[i][0] < pivot[0] {
			i++
		}
		for i < j && intervals[j][0] >= pivot[0] {
			j--
		}

		intervals[i], intervals[j] = intervals[j], intervals[i]
	}

	intervals[i], intervals[right] = intervals[right], intervals[i]
	return i
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
