package leetcode

// simple solution
func isValidSudokuSimply(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				index := board[i][j] - '0'
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}

	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				index := board[j][i] - '0'
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := [10]int{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					if board[ii][jj] != '.' {
						index := board[ii][jj] - '0'
						if tmp[index] == 1 {
							return false
						}
						tmp[index] = 1
					}
				}
			}
		}
	}

	return true
}

// hashMap solution
func isValidSudoku(board [][]byte) bool {
	// 分别缓存行、列、3×3宫格
	rows, cols, boxes := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				// 确定是第几个3*3的小区域
				index := i/3*3 + j/3
				if rows[i][num] || cols[j][num] || boxes[index][num] {
					return false
				}
				rows[i][num] = true
				cols[j][num] = true
				boxes[index][num] = true
			}
		}
	}

	return true
}
