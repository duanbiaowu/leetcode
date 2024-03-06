package leetcode

func sortedSquares(nums []int) []int {
	// 数组有序，但是负数平方可能大于正数平方
	// 例如 [-10, 1, 2, 5] 最小的负数平方之后大于最大的正数
	// 所以新数组从后往前插入元素
	// 左右双指针，根据平方后数值大小选择前进 OR 后退
	// 每次比较左右两个数，即使出现负数也 OK, 因为负数越小，其平方值越大
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1

	for index := right; index >= 0; index-- {
		// 这里有个优化小技巧:
		//   其实不需要对比两个数的平方，因为数组本身是已经排序完成的
		//   所以只需要将左边的数字取负数，即可验证两个数字平方后的大小
		//   举个例子: 左边数字 = -9 或者 9, 右边数字 = 10
		//   - (-9) = 9 < 10
		//   - (9) = -9 < 10
		//   PS: 由此可见如果左边的数字平方后比右边数字平方后小，其取负数依然比右边数字小
		if x, y := nums[left]*nums[left], nums[right]*nums[right]; x > y {
			res[index] = x
			left++
		} else {
			res[index] = y
			right--
		}
	}
	return res
}
