package leetcode

func bestLine(points [][]int) []int {
	var res []int
	num, maxNum, size := 0, 0, len(points)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size-1; j++ {
			num = 2
			// 剪枝
			if size-1-j+2 <= maxNum {
				break
			}
			x1 := points[j][0] - points[i][0]
			y1 := points[j][1] - points[i][1]

			for k := j + 1; k < size; k++ {
				x2 := points[k][0] - points[i][0]
				y2 := points[k][1] - points[i][1]
				// 满足向量共线的条件，且有一个公共顶点，所以在一条直线上
				if x1*y2 == x2*y1 {
					num++
				}
			}

			if num > maxNum {
				maxNum = num
				res = []int{i, j}
			}
		}
	}

	return res
}
