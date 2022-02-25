package leetcode

// 滑动窗口
func findContinuousSequence(target int) [][]int {
	// 注意初始值
	i, j, sum := 1, 2, 3
	var res [][]int
	for i < j {
		if sum == target {
			row := make([]int, j-i+1)
			for k := i; k <= j; k++ {
				row[k-i] = k
			}
			res = append(res, row)
		}
		if sum >= target {
			sum -= i
			i++ // 关键顺序
		} else {
			j++ // 关键顺序
			sum += j
		}
	}
	return res
}
