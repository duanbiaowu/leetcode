package leetcode

func minOperations(nums []int, k int) int {
	set := make(map[int]struct{})

	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] <= k {
			if _, ok := set[nums[i]]; !ok {
				set[nums[i]] = struct{}{}
			}
			if len(set) >= k {
				return n - i
			}
		}
	}

	return -1
}
