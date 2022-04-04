package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := make(map[int]int)
	for i := range nums1 {
		m[nums1[i]]++
	}

	var res []int
	for i := range nums2 {
		if cnt, ok := m[nums2[i]]; ok && cnt > 0 {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}
	return res
}
