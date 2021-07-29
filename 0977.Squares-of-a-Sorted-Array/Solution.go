package leetcode

func sortedSquares(nums []int) []int {
	// 数组有序，但是负数平方可能大于正数平方
	// 例如 [-10, 1, 2, 5] 最小的负数平方之后大于最大的正数
	// 所以新数组从后往前插入元素
	// 头尾双指针，根据平方后数值大小选择前进 OR 后退
	n := len(nums)
	res := make([]int, n)
	tail, head := 0, n-1

	for index := head; index >= 0; index-- {
		// golang syntactic sugar
		if x, y := nums[tail]*nums[tail], nums[head]*nums[head]; x > y {
			res[index] = x
			tail++
		} else {
			res[index] = y
			head--
		}
	}
	return res
}
