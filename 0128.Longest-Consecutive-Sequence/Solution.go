package leetcode

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
		if _, ok := set[v-1]; !ok {
			cnt := 1
			for {
				if _, ok = set[v+1]; ok {
					v++
					cnt++
				} else {
					break
				}
			}
			if cnt > longest {
				longest = cnt
			}
		}
	}
	return longest
}
