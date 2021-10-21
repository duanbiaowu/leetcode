package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	var (
		rows, columns            = len(matrix), len(matrix[0])
		res                      = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		// 方向: ->
		for column := left; column <= right; column++ {
			res[index] = matrix[top][column]
			index++
		}
		// 方向: ↓
		for row := top + 1; row <= bottom; row++ {
			res[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			// 方向: <-
			for column := right - 1; column > left; column-- {
				res[index] = matrix[bottom][column]
				index++
			}
			// 方向: ↑
			for row := bottom; row > top; row-- {
				res[index] = matrix[row][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return res
}
