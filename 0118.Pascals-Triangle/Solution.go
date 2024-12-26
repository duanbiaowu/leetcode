package leetcode

func generate(numRows int) [][]int {
	if numRows < 1 {
		return [][]int{}
	}

	res := make([][]int, numRows)

	for row := 0; row < numRows; row++ {
		// 初始化每一行的数组
		res[row] = make([]int, row+1)
		// 第一列和最后一列都是1
		res[row][0] = 1
		res[row][row] = 1

		// 计算当前行从 第 2 列到倒数第 2 列 的值
		for col := 1; col < row; col++ {
			res[row][col] = res[row-1][col-1] + res[row-1][col]
		}
	}

	return res
}
