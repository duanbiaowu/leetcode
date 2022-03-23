package leetcode

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return intervals
	}

	quickSort(intervals, 0, n-1)
	res := [][]int{intervals[0]}
	pos := 0

	for i := 1; i < n; i++ {
		if intervals[i][0] > res[pos][1] {
			res = append(res, intervals[i])
			pos++
		} else {
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
		for i < j && (intervals[i][0] < pivot[0] || intervals[i][0] == pivot[0] && intervals[i][1] < pivot[1]) {
			i++
		}
		// 非常细微的边界：intervals[j][1] >= pivot[1]
		for i < j && (intervals[j][0] > pivot[0] || intervals[j][0] == pivot[0] && intervals[j][1] >= pivot[1]) {
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
