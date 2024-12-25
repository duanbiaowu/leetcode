package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return false
}
