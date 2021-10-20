package leetcode

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			matrix[top][column] = num
			num++
		}
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				matrix[bottom][column] = num
				num++
			}
			for row := bottom; row > top; row-- {
				matrix[row][left] = num
				num++
			}
		}
		left++
		right--
		top++
		bottom--
	}

	return matrix
}
