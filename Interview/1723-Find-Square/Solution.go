package leetcode

func findSquare(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	//dp[i][j][0]: i,j左边连续的0的个数
	//dp[i][j][1]: i,j上边连续的0的个数
	dp := make([][][2]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][2]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if j > 0 {
					dp[i][j][0] = dp[i][j-1][0] + 1
				} else {
					dp[i][j][0] = 1
				}
				if i > 0 {
					dp[i][j][1] = dp[i-1][j][1] + 1
				} else {
					dp[i][j][1] = 1
				}
			}
		}
	}

	r, c, size := -1, -1, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 枚举以 [i,j] 为右下角的正方形的可能边长side
			for side := min(dp[i][j][0], dp[i][j][1]); side >= 1; side-- {
				if dp[i][j-side+1][1] >= side && dp[i-side+1][j][0] >= side {
					if side > size {
						size = side
						r = i - side + 1
						c = j - side + 1
						// 因为是从右下角开始，所以更小的正方形无需考虑了
						break
					}
				}
			}
		}
	}

	if r == -1 && c == -1 {
		return []int{}
	}
	return []int{r, c, size}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
