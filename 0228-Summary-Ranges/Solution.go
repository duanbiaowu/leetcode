package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	var res []string

	for right, n := 0, len(nums); right < n; {
		left := right
		right++

		for right < n && nums[right] == nums[right-1]+1 {
			right++
		}

		s := strconv.Itoa(nums[left])
		if left < right-1 {
			s += "->" + strconv.Itoa(nums[right-1])
		}

		res = append(res, s)
	}

	return res
}
