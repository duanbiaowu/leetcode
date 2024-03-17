package leetcode

func searchInsert(nums []int, target int) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] >= target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func searchInsertWithComment(nums []int, target int) int {
	low, hi := 0, len(nums)-1

	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] == target {
			// 针对不存在重复元素的情况
			// 找到目标值后直接返回索引
			return mid
		} else if nums[mid] >= target {
			// target 在数组的左半部分
			hi = mid - 1
		} else {
			// target 在数组的右半部分
			low = mid + 1
		}
	}

	// 如果没有找到目标值，直接将左边界返回
	// 因为此时的数组以 low 位置切分为两部分：
	// 1. 左半部分的元素全部比 nums[low] 小
	// 2. 右半部分的元素全部比 nums[low] 大
	// 索引 low 正好就是插入 target 目标元素的位置
	return low
}
