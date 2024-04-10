package leetcode

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	// dp[i][j]: 以 i 行 j 列为右下角所能构成的最大正方形边长
	// dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	maxSide := 0

	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
				if maxSide < dp[i][j] {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

// 暴力法 (超时)
// 超时原因: 重复检测
func maximalSquareSample(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	var res int
	rows, cols := len(matrix), len(matrix[0])

	// 按照矩阵逐个元素遍历
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			// 以当前坐标作为矩形左上角，检测最大正方形的边长
			// 更新已知的最大正方形面积
			res = max(res, maxArea(matrix, i, j))
		}
	}

	return res
}

func maxArea(matrix [][]byte, x, y int) int {
	rows, cols := len(matrix), len(matrix[0])

	curLength := 1

	// 边长从 2 开始递增
	for length := 2; x+length <= rows && y+length <= cols; length++ {
		for i := 0; i < length; i++ {
			// 如果当前 “新的正方形” 中右侧的元素包含 0
			// 直接返回对应的面积
			if matrix[x+i][y+curLength] == '0' {
				return curLength * curLength
			}
		}
		for i := 0; i < length; i++ {
			// 如果当前 “新的正方形” 中底侧的元素包含 0
			// 直接返回对应的面积
			if matrix[x+curLength][y+i] == '0' {
				return curLength * curLength
			}
		}

		curLength = length
	}

	return curLength * curLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	}
	if y < z {
		return y
	}
	return z
}
