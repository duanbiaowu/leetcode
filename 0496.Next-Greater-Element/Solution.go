package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))

	n := len(nums2)
	for i := range nums1 {
		res[i] = -1
		j := 0
		for j < n && nums2[j] != nums1[i] {
			j++
		}

		for ; j < n; j++ {
			if nums2[j] > nums1[i] {
				res[i] = nums2[j]
				break
			}
		}
	}

	return res
}

// reference: https://leetcode.cn/problems/next-greater-element-i/solutions/1065517/xia-yi-ge-geng-da-yuan-su-i-by-leetcode-bfcoj/
func nextGreaterElementOpt(nums1 []int, nums2 []int) []int {
	index := make(map[int]int)
	// 单调栈，用于存储中间状态值
	stack := []int{}

	// 预处理 num2 数组
	for i := len(nums2) - 1; i >= 0; i-- {
		// 只要当前数字比栈顶部数字要大，说明当前数字右侧不存在比当前数字更大的数字了
		// 不断出栈顶部数字即可
		for len(stack) > 0 && nums2[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			// 找到了右侧第一个比当前数字更大的数字，更新索引
			index[nums2[i]] = stack[len(stack)-1]
		} else {
			// 说明右侧不存在比当前数字更大的数字了，更新索引为 -1
			index[nums2[i]] = -1
		}
		// 将当前值状态中间值入栈
		stack = append(stack, nums2[i])
	}

	// 遍历 num1 数组，填充索引值即可
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = index[num]
	}

	return res
}
