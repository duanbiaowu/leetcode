package leetcode

import "sort"

// 按照 峰-谷-峰-谷 排列
// 奇数位置 = 峰值，判断是否比前一个元素大
// 偶数位置 = 谷值，判断是否比前一个元素小
func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if i%2 == 1 {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
}

func wiggleSort2(nums []int) {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
}
