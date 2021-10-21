package leetcode

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	quickSort(intervals, 0, len(intervals)-1)
	res := [][]int{intervals[0]}
	index := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[index][1] {
			index++
			res = append(res, intervals[i])
		} else {
			res[index][1] = max(res[index][1], intervals[i][1])
		}
	}

	return res
}

func partition(a [][]int, left, right int) int {
	pivot := a[right]
	for i := left; i < right; i++ {
		if (a[i][0] < pivot[0]) || (a[i][0] == pivot[0] && a[i][1] < pivot[1]) {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	return left
}

func quickSort(a [][]int, left, right int) {
	if left < right {
		p := partition(a, left, right)
		quickSort(a, left, p-1)
		quickSort(a, p+1, right)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
