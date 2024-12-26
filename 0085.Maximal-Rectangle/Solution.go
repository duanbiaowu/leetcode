package leetcode

import leetcode "leetcode/0084.Largest-Rectangle-in-Histogram"

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	res := 0

	// Example:
	// {'1', '0', '1', '0', '0'}
	// {'1', '0', '1', '1', '1'}
	// {'1', '1', '1', '1', '1'}
	// {'1', '0', '0', '1', '0'}

	// heights init = [0 0 0 0 0]
	// heights = [1 0 1 0 0]
	// heights = [2 0 2 1 1]
	// heights = [3 1 3 2 2]
	// heights = [4 0 0 3 0]

	heights := make([]int, len(matrix[0]))
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '0' {
				heights[col] = 0
			} else {
				heights[col]++
			}
		}

		// 对于每一列，使用基于柱状图的方法
		res = max(res, leetcode.LargestRectangleArea(heights))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
