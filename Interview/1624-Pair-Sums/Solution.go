package leetcode

import "sort"

func pairSums(nums []int, target int) [][]int {
	var res [][]int
	m := make(map[int]int)
	for i := range nums {
		if val, ok := m[target-nums[i]]; ok && val > 0 {
			res = append(res, []int{target - nums[i], nums[i]})
			m[target-nums[i]]--
		} else {
			m[nums[i]]++
		}
	}
	return res
}

func pairSums2(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			res = append(res, []int{nums[left], nums[right]})
			left++
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}
	return res
}
