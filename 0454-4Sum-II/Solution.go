package leetcode

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)

	for i := range nums1 {
		for j := range nums2 {
			sum := nums1[i] + nums2[j]
			if _, ok := m[sum]; ok {
				m[sum]++
			} else {
				m[sum] = 1
			}
		}
	}

	cnt := 0
	for i := range nums3 {
		for j := range nums4 {
			sum := 0 - nums3[i] - nums4[j]
			if val, ok := m[sum]; ok {
				cnt += val
			}
		}
	}

	return cnt
}
