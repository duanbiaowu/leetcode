package leetcode

// 将所有正数作为数组下标，置对应数组值为负值。
// 那么，仍为正数的位置即为（未出现过）消失的数字。
// 	原始数组：[4,3,2,7,8,2,3,1]
// 	重置后为：[-4,-3,-2,-7,8,2,-3,-1]
// 	[8,2] 分别对应的index为[5,6]
func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		p := abs(nums[i]) - 1
		nums[p] = -abs(nums[p])
	}

	var res []int
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
