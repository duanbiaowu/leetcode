package leetcode

// 将矩阵每一行拼接在上一行的末尾，则会得到一个升序数组
// 可以在该数组上二分查找
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	// 边界处理
	if rows == 0 {
		return false
	}

	cols := len(matrix[0])
	// 一维数组的长度等于矩阵的行数量 * 列数量 - 1
	// 因为数组下标从 0 开始
	low, hi := 0, rows*cols-1

	for low <= hi {
		mid := low + (hi-low)>>1
		// 将索引转换为具体的矩阵坐标
		// mid / cols 得到 X 坐标
		// mid % cols 得到 Y 坐标
		val := matrix[mid/cols][mid%cols]

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
