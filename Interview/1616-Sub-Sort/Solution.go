package leetcode

import "math"

// 题目的核心是要找出 “有问题” 的中间段
// 分别求左右两边 - 分治
func subSort(array []int) []int {
	start, end := -1, -1
	min, max := math.MaxInt32, math.MinInt32

	// 从前往后找目标末位，使得从该位到最后，数组保持递增
	for i := 0; i < len(array); i++ {
		if array[i] >= max {
			max = array[i]
		} else {
			end = i
		}
	}

	// 数组恒递增，说明数组是有序的，直接返回
	if end == -1 {
		return []int{start, end}
	}

	// 从后往前找目标首位，使得从该位到最前，数组保持递减
	for i := end; i >= 0; i-- {
		if array[i] <= min {
			min = array[i]
		} else {
			start = i
		}
	}

	return []int{start, end}
}
