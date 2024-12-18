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
				// 如果当前坐标的值为 1
				// 说明正方形的边长至少为 1，所以 + 1
				// 在此基础上分别获取以 左上角、上边、左边 3 个左边为右下角坐标形成的正方形数量，并取其中的最小值 + 1
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}

// 暴力搜索迭代版本 (超时)
// 超时原因: 重复检测
func maximalSquareBruteForce(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	var maxSide int
	rows, cols := len(matrix), len(matrix[0])

	// 按照矩阵逐个元素遍历
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			// 以当前坐标作为矩形左上角，检测最大正方形的边长
			// 更新已知的最大正方形面积
			maxSide = max(maxSide, search(matrix, i, j))
		}
	}

	return maxSide * maxSide
}

// 搜索以当前坐标 [x, y] 为左上角的最大正方形边长
func search(matrix [][]byte, x, y int) int {
	rows, cols := len(matrix), len(matrix[0])

	curSide := 1

	// 边长从 2 开始递增
	for length := 2; x+length <= rows && y+length <= cols; length++ {
		for i := 0; i < length; i++ {
			// 如果当前 “新的正方形” 中右侧的元素包含 0
			// 直接返回对应的边长
			if matrix[x+i][y+curSide] == '0' {
				return curSide
			}
		}
		for i := 0; i < length; i++ {
			// 如果当前 “新的正方形” 中底侧的元素包含 0
			// 直接返回对应的边长
			if matrix[x+curSide][y+i] == '0' {
				return curSide
			}
		}

		curSide = length
	}

	return curSide
}

// 暴力搜索递归版本 (超时)
// 超时原因: 重复检测
func maximalSquareBruteForceRecursive(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	maxSide := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxSide = max(maxSide, searchRecursive(matrix, i, j))
		}
	}

	return maxSide * maxSide
}

// 搜索以当前坐标 [x, y] 为左上角的最大正方形边长
func searchRecursive(matrix [][]byte, x, y int) int {
	if x >= len(matrix) || y >= len(matrix[0]) {
		return 0
	}
	if matrix[x][y] == '0' {
		return 0
	}

	// 搜索以 [当前坐标的底侧元素] 为坐标 的最大正常形边长
	down := searchRecursive(matrix, x+1, y)
	// 搜索以 [当前坐标的右侧元素] 为坐标 的最大正常形边长
	right := searchRecursive(matrix, x, y+1)
	// 搜索以 [当前坐标的右下侧元素] 为坐标 的最大正常形边长
	diag := searchRecursive(matrix, x+1, y+1)

	return 1 + min(down, right, diag)
}

// 记忆化搜索
func maximalSquareMemo(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	maxSide := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxSide = max(maxSide, searchMemo(matrix, i, j, memo))
		}
	}

	return maxSide * maxSide
}

func searchMemo(matrix [][]byte, x, y int, memo [][]int) int {
	if x >= len(matrix) || y >= len(matrix[0]) {
		return 0
	}
	if matrix[x][y] == '0' {
		return 0
	}

	// 如果备忘录中已经包含 [当前坐标] 的边长
	// 直接返回
	if memo[x][y] != 0 {
		return memo[x][y]
	}

	// 搜索以 [当前坐标的底侧元素] 为坐标 的最大正常形边长
	down := searchMemo(matrix, x+1, y, memo)
	// 搜索以 [当前坐标的右侧元素] 为坐标 的最大正常形边长
	right := searchMemo(matrix, x, y+1, memo)
	// 搜索以 [当前坐标的右下侧元素] 为坐标 的最大正常形边长
	diag := searchMemo(matrix, x+1, y+1, memo)

	// 更新备忘录中 [当前坐标] 的边长
	memo[x][y] = 1 + min(down, right, diag)

	return memo[x][y]
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
