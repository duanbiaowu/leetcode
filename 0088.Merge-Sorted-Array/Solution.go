package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	right := m + n - 1
	p1, p2 := m-1, n-1

	for ; p1 >= 0 && p2 >= 0; right-- {
		if nums1[p1] > nums2[p2] {
			nums1[right] = nums1[p1]
			p1--
		} else {
			nums1[right] = nums2[p2]
			p2--
		}
	}

	// 题目假设 nums1 的空间大小等于 m + n
	// 这样它就有足够的空间保存来自 nums2 的元素
	// 所以这里 p2 存在两种情况
	// 	(1) >  0 说明两个数组中最小的元素在 nums1，无需调整
	//  (2) == 0 说明两个数组中最小的元素在 nums2，仅调整 nums2 中剩余元素
	for p2 >= 0 {
		nums1[right] = nums2[p2]
		right--
		p2--
	}
}

// 参考归并排序中的 “合并过程” 代码，但是会造成额外的空间消耗
func merge2(nums1 []int, m int, nums2 []int, n int) {
	// 定义数组 (切片时) 直接声明好长度，避免扩容
	sorted := make([]int, m+n)
	// 用于更新已排序数组中元素的索引
	index := 0

	p1, p2 := 0, 0
	for ; p1 < m && p2 < n; index++ {
		if nums1[p1] < nums2[p2] {
			sorted[index] = nums1[p1]
			p1++
		} else {
			sorted[index] = nums2[p2]
			p2++
		}
	}

	// 将 nums1 数组中剩余元素未添加到已排序数组
	for ; p1 < m; index++ {
		sorted[index] = nums1[p1]
		p1++
	}
	// 将 nums2 数组中剩余元素未添加到已排序数组
	for ; p2 < n; index++ {
		sorted[index] = nums2[p2]
		p2++
	}

	// 将合并完成的数组赋值到 nums1
	for i := range nums1 {
		nums1[i] = sorted[i]
	}
}
