package leetcode

import "sort"

// 第一时间想到的是 Sort, 然后计算最长序列
// 但题目要求时间复杂度为：O(N)
// 空间换时间
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for i := range nums {
		set[nums[i]] = struct{}{}
	}

	longest := 0

	for v := range set {
		// 如果当前值 -1 存在于集合中
		// 那么当前值必然已经作为其前一个值的 “连续序列” 计算过了
		// 跳过即可
		if _, ok := set[v-1]; ok {
			continue
		}

		cnt := 1
		// 计算以当前值为基准开始的最长序列
		for {
			if _, ok := set[v+1]; ok {
				cnt++
				v++
			} else {
				break
			}
		}

		// 更新最长序列长度
		if cnt > longest {
			longest = cnt
		}
	}

	return longest
}

func longestConsecutive2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	sort.Ints(nums)

	maxLen := 1
	curLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			curLen++
		} else {
			curLen = 1
		}

		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
