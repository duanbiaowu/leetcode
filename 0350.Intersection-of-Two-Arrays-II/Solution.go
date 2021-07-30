package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := map[int]int{}
	for _, val := range nums1 {
		m[val]++
	}

	res := []int{}
	for _, val := range nums2 {
		if count, ok := m[val]; ok && count > 0 {
			res = append(res, val)
			m[val]--
		}
	}
	return res
}
