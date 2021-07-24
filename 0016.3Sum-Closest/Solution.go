package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n   = len(nums)
		res = math.MaxInt32
	)

	// 两个小的优化
	// 1. a+b+c == target 直接返回
	// 2. 枚举 first, second, third 中任意元素并移动指针时
	//    可以直接将其移动到下一个不相同的元素，减少枚举的次数
	//	  ( 因为求的是最接近的值, 并不是元素组合, 所以相同的元素计算一次即可 )
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second, third := first+1, n-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum > target {
				third--
				// 移到左边第一个不相等元素
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			} else {
				second++
				// 移到右边第一个不相等元素
				for second < third && nums[second] == nums[second-1] {
					second++
				}
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
