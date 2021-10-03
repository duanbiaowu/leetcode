package leetcode

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	res := 0
	// 对于每一列，使用基于柱状图的方法
	for j := 0; j < n; j++ {
		up := make([]int, m)
		down := make([]int, m)
		var stack []int

		for i := 0; i < m; i++ {
			for len(stack) > 0 && left[stack[len(stack)-1]][j] >= left[i][j] {
				stack = stack[:len(stack)-1]
			}
			up[i] = -1
			if len(stack) > 0 {
				up[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}

		stack = nil
		for i := m - 1; i >= 0; i-- {
			for len(stack) > 0 && left[stack[len(stack)-1]][j] >= left[i][j] {
				stack = stack[:len(stack)-1]
			}
			down[i] = m
			if len(stack) > 0 {
				down[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}

		for i := 0; i < m; i++ {
			res = max(res, (down[i]-up[i]-1)*left[i][j])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
