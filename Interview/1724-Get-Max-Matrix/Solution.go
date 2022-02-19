package leetcode

func getMaxMatrix(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	max := matrix[0][0]
	res := []int{0, 0, 0, 0}

	// 构造列的前缀和
	preSum := make([][]int, m+1)
	preSum[0] = make([]int, n)
	for i := 1; i <= m; i++ {
		preSum[i] = make([]int, n)
		for j := 0; j < n; j++ {
			preSum[i][j] = preSum[i-1][j] + matrix[i-1][j]
		}
	}

	// 合并行
	for top := 0; top < m; top++ {
		for bottom := top; bottom < m; bottom++ {
			// 构造一维矩阵
			row := make([]int, n)
			for i := 0; i < n; i++ {
				row[i] = preSum[bottom+1][i] - preSum[top][i]
			}
			// 最大子数组问题
			start := 0
			sum := row[0]
			for i := 1; i < n; i++ {
				if sum > 0 {
					sum += row[i]
				} else {
					sum = row[i]
					start = i
				}
				if sum > max {
					max = sum
					res[0] = top
					res[1] = start
					res[2] = bottom
					res[3] = i
				}
			}
		}
	}
	return res
}
