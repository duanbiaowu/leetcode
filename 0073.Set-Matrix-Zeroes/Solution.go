package leetcode

func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])

	// 记录第一行和第一列是否原本包含 0
	row0, col0 := false, false
	for _, val := range matrix[0] {
		if val == 0 {
			row0 = true
			break
		}
	}
	for _, row := range matrix {
		if row[0] == 0 {
			col0 = true
			break
		}
	}

	// 使用其他行与列更新 第一行与第一列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 使用第一行与第一列 更新其他行与列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 更新第一行和第一列是否原本包含 0
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, row := range matrix {
			row[0] = 0
		}
	}
}

// 一个简单的改进方案是使用 O(m + n) 的额外空间
// 但这仍然不是最好的解决方案
func setZeroesSimple(matrix [][]int) {
	// 边界条件处理
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	rowMap := make(map[int]struct{})
	colMap := make(map[int]struct{})

	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				rowMap[i] = struct{}{}
				colMap[j] = struct{}{}
			}
		}
	}

	for i, row := range matrix {
		for j := range row {
			if _, ok := rowMap[i]; ok {
				row[j] = 0
			} else if _, ok := colMap[j]; ok {
				row[j] = 0
			}
		}
	}
}
