package leetcode

func containsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
