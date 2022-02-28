package leetcode

import "sort"

func isStraight(nums []int) bool {
	set := make(map[int]struct{})
	// A-K + 大小王（算 1 张） = 14
	max, min := 0, 14
	for i := range nums {
		// 大小王可当作任何牌
		if nums[i] == 0 {
			continue
		}
		// 重复牌
		if _, ok := set[nums[i]]; ok {
			return false
		}
		set[nums[i]] = struct{}{}
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return max-min < 5
}

// 排序
func isStraight2(nums []int) bool {
	joker := 0
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	// nums[joker]: 很巧妙，大小王可以当作几张牌
	return nums[4]-nums[joker] < 5
}
