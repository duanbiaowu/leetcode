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

func isValidSudokuSimple2(board [][]byte) bool {
	if len(board) == 0 || len(board)%9 != 0 || len(board[0])%9 != 0 {
		return false
	}

	rows := len(board)
	cols := len(board[0])

	// check every row
	for _, row := range board {
		m := make(map[byte]struct{})
		for _, c := range row {
			if c == '.' {
				continue
			}
			if _, ok := m[c]; ok {
				return false
			}
			m[c] = struct{}{}
		}
	}

	// check every column
	for i := 0; i < rows; i++ {
		m := make(map[byte]struct{})
		for j := 0; j < cols; j++ {
			if board[j][i] == '.' {
				continue
			}
			if _, ok := m[board[j][i]]; ok {
				return false
			}
			m[board[j][i]] = struct{}{}
		}
	}

	// check every 3*3
	for i := 0; i < rows/3; i++ {
		for j := 0; j < cols/3; j++ {
			m := make(map[byte]struct{})

			for k := i * 3; k < i*3+3; k++ {
				for k2 := j * 3; k2 < j*3+3; k2++ {
					if board[k][k2] == '.' {
						continue
					}
					if _, ok := m[board[k][k2]]; ok {
						return false
					}
					m[board[k][k2]] = struct{}{}
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
