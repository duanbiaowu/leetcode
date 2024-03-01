package leetcode

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var res []int
	mask := make(map[int]int)

	// 第一个数组中的元素不管出现多少次，赋值为 1 作为存在标志位
	for _, num := range nums1 {
		if _, ok := mask[num]; !ok {
			mask[num] = 1
		}
	}

	// 第二个数组中的元素
	//   如果已经出现在第一个数组中，追加到结果中，并赋值为 -1 作为标志位表示多次重复出现
	//   如果是首次出现，并赋值为 2 作为存在标志位
	for _, num := range nums2 {
		if _, ok := mask[num]; !ok {
			mask[num] = 2
		} else if mask[num] == 1 {
			mask[num] = -1
			res = append(res, num)
		}
	}

	// 第三个数组的元素
	//   如果已经出现在第一个或第二个数组中，追加到结果中，并赋值为 -1 作为标志位表示多次重复出现
	//   其他情况不予考虑 (重复)
	for _, num := range nums3 {
		if v, ok := mask[num]; ok {
			if v == 1 || v == 2 {
				mask[num] = -1
				res = append(res, num)
			}
		}
	}

	return res
}

// 参数有 3 个数组，使用 Mask 中元素的最低 3 个二进制位来表示数组元素位于哪个数组中
// 例如 0000 0101 表示数字出现在第一个和第三个数组中
// 最后判断二进制表示中位为 1 的个数即可
//
//	判断条件为大于 1, 但是不需要知道具体的位为 1 的数量，因此可以简写为 v&(v-1) > 0, 只要大于 0, 说明数量至少在 1 个以上
func twoOutOfThreeOpt(nums1 []int, nums2 []int, nums3 []int) []int {
	mask := make(map[int]int)

	for _, num := range nums1 {
		mask[num] |= 1 << 0
	}
	for _, num := range nums2 {
		mask[num] |= 1 << 1
	}
	for _, num := range nums3 {
		mask[num] |= 1 << 2
	}

	var res []int
	for k, v := range mask {
		if v&(v-1) > 0 {
			res = append(res, k)
		}
	}

	return res
}
