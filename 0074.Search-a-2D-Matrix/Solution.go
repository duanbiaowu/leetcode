package leetcode

// 将矩阵每一行拼接在上一行的末尾，则会得到一个升序数组
// 可以在该数组上二分查找
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	low, hi := 0, m*n-1
	for low <= hi {
		mid := low + (hi-low)>>1
		val := matrix[mid/n][mid%n]
		if val == target {
			return true
		} else if val < target {
			low = mid + 1
		} else {
			hi = mid - 1
		}

	}

	return false
}
