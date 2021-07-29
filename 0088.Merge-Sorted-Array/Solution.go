package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	fast := m + n - 1
	slowN1, slowN2 := m-1, n-1
	for slowN1 >= 0 && slowN2 >= 0 {
		if nums1[slowN1] > nums2[slowN2] {
			nums1[fast] = nums1[slowN1]
			slowN1--
		} else {
			nums1[fast] = nums2[slowN2]
			slowN2--
		}
		fast--
	}

	// 题目假设 nums1 的空间大小等于 m + n
	// 这样它就有足够的空间保存来自 nums2 的元素
	// 所以这里 slowN1 存在两种情况
	// 	(1) >  0 说明两个数组中最小的元素在 nums1，无需调整
	//  (2) == 0 说明两个数组中最小的元素在 nums2，仅调整 nums2 中剩余元素
	for slowN2 >= 0 {
		nums1[fast] = nums2[slowN2]
		fast--
		slowN2--
	}
}
