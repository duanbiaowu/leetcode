package leetcode

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])

	// 记录第一行和第一列是否原本包含 0
	firstRowHasZero, firstColHasZero := false, false
	for _, col := range matrix[0] {
		if col == 0 {
			firstRowHasZero = true
			break
		}
	}
	for _, row := range matrix {
		if row[0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// 使用其他行与列更新 第一行与第一列
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 使用第一行与第一列 更新其他行与列
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 更新第一行和第一列是否原本包含 0
	if firstRowHasZero {
		for i := 0; i < cols; i++ {
			matrix[0][i] = 0
		}
	}
	if firstColHasZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
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
