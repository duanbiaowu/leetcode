package leetcode

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	// 构造前缀和
	for i, row := range matrix {
		sums[i] = make([]int, len(row)+1)
		for j, v := range row {
			sums[i][j+1] = sums[i][j] + v
		}
	}
	return NumMatrix{sums}
}

// 直接计算前缀区间总和
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 < 0 || row1 >= len(this.sums) || row2 < 0 || row2 >= len(this.sums) ||
		col1 < 0 || col1 >= len(this.sums[0]) || col2 < 0 || col2 >= len(this.sums[0]) {
		return 0
	}

	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.sums[i][col2+1] - this.sums[i][col1]
	}
	return sum
}
