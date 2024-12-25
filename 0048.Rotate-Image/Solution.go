package leetcode

// -----------
//	1   2   3
//	4   5   6
//	7   8   9
// -----------

func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n>>1; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	// -----------
	//  7   8   9
	//  4   5   6
	//  1   2   3
	// -----------

	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// -----------
	//  7   4   1
	//  8   5   2
	//  9   6   3
	// -----------
}
